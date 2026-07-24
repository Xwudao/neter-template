package mdw

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/gin-gonic/gin"

	"github.com/Xwudao/neter-template/internal/domain/payloads"
)

type indexModifierFunc func([]byte, *gin.Context) (*payloads.SeoPayload, error)

func (f indexModifierFunc) SEO(html []byte, c *gin.Context) (*payloads.SeoPayload, error) {
	return f(html, c)
}

func testSpaFS() fs.FS {
	return fstest.MapFS{
		"dist/index.html":             &fstest.MapFile{Data: []byte("<html><body>app</body></html>")},
		"dist/.vite/manifest.json":    &fstest.MapFile{Data: []byte(`{"main.ts":{"file":"assets/main.js","isEntry":true}}`)},
		"dist/assets/main.js":         &fstest.MapFile{Data: []byte("console.log('app')")},
		"dist/assets/main.js.gz":      &fstest.MapFile{Data: []byte("compressed app")},
		"dist/assets/directory/.keep": &fstest.MapFile{Data: []byte("placeholder")},
	}
}

func newTestSpa(t *testing.T) *SpaMdw {
	t.Helper()
	spa, err := NewSpaMdw(testSpaFS(), "dist", indexModifierFunc(func(html []byte, _ *gin.Context) (*payloads.SeoPayload, error) {
		return &payloads.SeoPayload{Ret: append(html, "<!-- modified -->"...), StatusCode: http.StatusOK}, nil
	}))
	if err != nil {
		t.Fatal(err)
	}
	return spa
}

func TestSpaMdwServesEveryIndexPathThroughModifier(t *testing.T) {
	t.Parallel()

	spa := newTestSpa(t)
	router := gin.New()
	router.NoRoute(spa.Serve("/"))

	for _, requestPath := range []string{"/", "/index.html", "/assets/directory"} {
		t.Run(requestPath, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, requestPath, nil)
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)

			if response.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
			}
			if !strings.Contains(response.Body.String(), "<!-- modified -->") {
				t.Fatal("SPA index did not pass through the modifier")
			}
			if cacheControl := response.Header().Get("Cache-Control"); cacheControl != "no-cache, no-store, must-revalidate" {
				t.Fatalf("Cache-Control = %q, want no-cache, no-store, must-revalidate", cacheControl)
			}
		})
	}
}

func TestSpaMdwStaticFilesStayInsideStaticRoot(t *testing.T) {
	t.Parallel()

	staticDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(staticDir, "readme.txt"), []byte("hello static"), 0o600); err != nil {
		t.Fatal(err)
	}

	spa := newTestSpa(t)
	spa.staticDir = staticDir
	router := gin.New()
	router.NoRoute(spa.Serve("/"))

	t.Run("regular static file supports ranges", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/readme.txt", nil)
		request.Header.Set("Range", "bytes=0-4")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusPartialContent {
			t.Fatalf("status = %d, want %d", response.Code, http.StatusPartialContent)
		}
		if got := response.Body.String(); got != "hello" {
			t.Fatalf("body = %q, want hello", got)
		}
	})

	t.Run("traversal request is rejected", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.URL.Path = "/../../go.mod"
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusNotFound {
			t.Fatalf("status = %d, want %d", response.Code, http.StatusNotFound)
		}
	})
}

func TestSpaMdwPrecompressedAssetsRespectAcceptEncoding(t *testing.T) {
	t.Parallel()

	spa := newTestSpa(t)
	router := gin.New()
	router.NoRoute(spa.Serve("/"))

	for _, tc := range []struct {
		name           string
		acceptEncoding string
		wantEncoding   string
	}{
		{name: "gzip", acceptEncoding: "br, gzip", wantEncoding: "gzip"},
		{name: "gzip disabled", acceptEncoding: "gzip;q=0", wantEncoding: ""},
		{name: "wildcard", acceptEncoding: "*;q=0.5", wantEncoding: "gzip"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/assets/main.js", nil)
			request.Header.Set("Accept-Encoding", tc.acceptEncoding)
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)

			if response.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
			}
			if got := response.Header().Get("Content-Encoding"); got != tc.wantEncoding {
				t.Fatalf("Content-Encoding = %q, want %q", got, tc.wantEncoding)
			}
		})
	}

	t.Run("compressed sidecar is not served directly", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/assets/main.js.gz", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusNotFound {
			t.Fatalf("status = %d, want %d", response.Code, http.StatusNotFound)
		}
	})
}

func TestSpaRelativePath(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		path   string
		prefix string
		want   string
		ok     bool
	}{
		{path: "/search", prefix: "/", want: "search", ok: true},
		{path: "/history/search", prefix: "/history/", want: "search", ok: true},
		{path: "/history/", prefix: "/history/", want: "", ok: true},
		{path: "/../../go.mod", prefix: "/", ok: false},
		{path: "/assets//index.js", prefix: "/", ok: false},
		{path: "/not-history", prefix: "/history/", ok: false},
	} {
		t.Run(tc.path, func(t *testing.T) {
			got, ok := spaRelativePath(tc.path, tc.prefix)
			if got != tc.want || ok != tc.ok {
				t.Fatalf("spaRelativePath(%q, %q) = (%q, %t), want (%q, %t)", tc.path, tc.prefix, got, ok, tc.want, tc.ok)
			}
		})
	}
}
