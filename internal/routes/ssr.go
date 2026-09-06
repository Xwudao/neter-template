package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"regexp"
	"sort"
	"strings"

	gossr "github.com/Xwudao/gotossr"
	"github.com/gin-gonic/gin"

	"github.com/Xwudao/neter-template/assets"
)

var (
	moduleScript = regexp.MustCompile(`(?s)<script type="module">.*?</script>`)
	headClose    = regexp.MustCompile(`</head>`)
)

type viteManifestEntry struct {
	CSS  []string `json:"css"`
	File string   `json:"file"`
}

// SSRPageData is the server-owned data contract for an SSR response. Replace
// ssrPageDataForPath with a business-layer query when pages are backed by the
// database; the result is used for both document metadata and SSR props.
type SSRPageData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
}

func ssrPageDataForPath(requestPath string) SSRPageData {
	pages := map[string]SSRPageData{
		"/": {
			Title:       "首页｜neter-template",
			Description: "测试首页描述：现代全栈模板，提供 React、TanStack Router 与 Go SSR。",
			Keywords:    "neter-template,React,TanStack Router,Go,SSR",
		},
		"/latest": {
			Title:       "最新资源｜neter-template",
			Description: "测试最新资源描述：浏览最近发布和更新的站内资源。",
			Keywords:    "最新资源,资源列表,neter-template",
		},
		"/tags": {
			Title:       "标签浏览｜neter-template",
			Description: "测试标签页描述：按主题筛选和浏览站内内容。",
			Keywords:    "标签,主题,内容分类,neter-template",
		},
		"/search": {
			Title:       "搜索｜neter-template",
			Description: "测试搜索页描述：搜索站内资源与文档。",
			Keywords:    "搜索,资源,文档,neter-template",
		},
		"/docs": {
			Title:       "文档｜neter-template",
			Description: "测试文档页描述：了解 neter-template 的使用方式和 SSR 集成。",
			Keywords:    "文档,SSR,Go,React,neter-template",
		},
		"/about": {
			Title:       "关于｜neter-template",
			Description: "测试关于页描述：neter-template 是一个现代化全栈起步模板。",
			Keywords:    "关于,全栈模板,neter-template",
		},
	}
	if page, ok := pages[requestPath]; ok {
		return page
	}
	return SSRPageData{
		Title:       "页面不存在｜neter-template",
		Description: "测试默认描述：你访问的页面不存在或已被移除。",
		Keywords:    "404,neter-template",
	}
}

// SSRRenderer renders public TanStack Router pages. Admin routes intentionally
// fall through to the Vite SPA shell so they remain client-only.
type SSRRenderer struct {
	engine     *gossr.Engine
	viteCSS    []byte
	viteScript []byte
}

func NewSSRRenderer() (*SSRRenderer, error) {
	viteCSS, viteScript, err := viteEntryTags()
	if err != nil {
		return nil, err
	}

	engine, err := gossr.New(gossr.Config{
		AppEnv:            "production",
		AssetRoute:        "/assets",
		FrontendDir:       "./web",
		ClientAppPath:     "src/ssr.tsx",
		SPAHydrationMode:  "tanstack",
		JSRuntimePoolSize: 1,
	})
	if err != nil {
		return nil, fmt.Errorf("create SSR engine: %w", err)
	}

	return &SSRRenderer{engine: engine, viteCSS: viteCSS, viteScript: viteScript}, nil
}

func (r *SSRRenderer) Shutdown(ctx context.Context) error {
	return r.engine.Shutdown(ctx)
}

func (r *SSRRenderer) Serve(spa gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isSSRRequest(c.Request) {
			spa(c)
			return
		}

		pageData := ssrPageDataForPath(c.Request.URL.Path)
		page := r.engine.RenderRoute(gossr.RenderConfig{
			File:  "src/ssr.tsx",
			Title: pageData.Title,
			MetaTags: map[string]string{
				"description": pageData.Description,
				"keywords":    pageData.Keywords,
			},
			Props:       pageData,
			RequestPath: c.Request.URL.RequestURI(),
		})
		// Stylesheets belong in <head>. Replacing the generated module script
		// directly would place them in <body>, which is what the page source
		// previously showed.
		page = headClose.ReplaceAll(page, append(r.viteCSS, []byte(`</head>`)...))
		page = moduleScript.ReplaceAll(page, r.viteScript)
		c.Data(http.StatusOK, "text/html; charset=utf-8", page)
	}
}

func viteEntryTags() ([]byte, []byte, error) {
	contents, err := fs.ReadFile(assets.SpaDist, "dist/.vite/manifest.json")
	if err != nil {
		return nil, nil, fmt.Errorf("read embedded Vite manifest: %w (run pnpm --dir web build first)", err)
	}

	var manifest map[string]viteManifestEntry
	if err := json.Unmarshal(contents, &manifest); err != nil {
		return nil, nil, fmt.Errorf("parse embedded Vite manifest: %w", err)
	}
	entry, ok := manifest["index.html"]
	if !ok || entry.File == "" {
		return nil, nil, fmt.Errorf("embedded Vite manifest has no index.html entry")
	}

	// TanStack Router code-splits route components. CSS Modules used by those
	// dynamic chunks are recorded on their own manifest entries, not on
	// index.html. A server-rendered document must include them up front:
	// otherwise the client can render the route before Vite injects its
	// chunk-local stylesheet (and crawlers never receive those styles at all).
	cssFiles := make(map[string]struct{})
	for _, manifestEntry := range manifest {
		for _, css := range manifestEntry.CSS {
			cssFiles[css] = struct{}{}
		}
	}

	cssPaths := make([]string, 0, len(cssFiles))
	for css := range cssFiles {
		cssPaths = append(cssPaths, css)
	}
	sort.Strings(cssPaths)

	var cssTags strings.Builder
	for _, css := range cssPaths {
		fmt.Fprintf(&cssTags, `<link rel="stylesheet" href="/%s">`, path.Clean(css))
	}
	scriptTag := fmt.Sprintf(`<script type="module" crossorigin src="/%s"></script>`, path.Clean(entry.File))
	return []byte(cssTags.String()), []byte(scriptTag), nil
}

func isSSRRequest(request *http.Request) bool {
	if request.Method != http.MethodGet && request.Method != http.MethodHead {
		return false
	}

	requestPath := request.URL.Path
	if requestPath == "/admin" || strings.HasPrefix(requestPath, "/admin/") {
		return false
	}
	if strings.HasPrefix(requestPath, "/api/") || strings.HasPrefix(requestPath, "/assets/") {
		return false
	}
	return path.Ext(requestPath) == ""
}
