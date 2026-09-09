package mdw

import (
	"io"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Xwudao/neter-template/internal/domain/payloads"
	"github.com/Xwudao/neter-template/pkg/libx"
)

const indexFile = "index.html"

type ServeFileSystem interface {
	http.FileSystem
}

// IndexModifier keeps SPA rendering testable without coupling the middleware to SeoBizBiz.
type IndexModifier interface {
	SEO(html []byte, c *gin.Context) (*payloads.SeoPayload, error)
}

type SpaMdw struct {
	fsData    fs.FS
	target    string
	staticDir string
	modifier  IndexModifier
}

func NewSpaMdw(fsData fs.FS, target string, modifier IndexModifier) (*SpaMdw, error) {
	workingDir, err := os.Getwd()
	staticDir := ""
	if err == nil {
		staticDir = filepath.Join(workingDir, "static")
	}

	sm := &SpaMdw{
		fsData:    fsData,
		target:    target,
		staticDir: staticDir,
		modifier:  modifier,
	}
	if err := sm.parseManifest(); err != nil {
		return nil, err
	}
	return sm, nil
}

// parseManifest validates the embedded Vite bundle at startup. A missing
// manifest is not fatal: the backend can run before the frontend has been
// built, and the SPA handler already degrades to 404 until assets exist.
func (m *SpaMdw) parseManifest() error {
	fd := EmbedFolder(m.fsData, m.target)
	manifest, err := readEmbedFile(fd, ".vite/manifest.json")
	if err != nil {
		log.Printf("warning: embedded SPA manifest not found (%v); run `nr run --web` or `nr build --web` to build the frontend", err)
		return nil
	}
	_, err = libx.ParseManifestString(string(manifest))
	return err
}

func (m *SpaMdw) Serve(urlPrefix string) gin.HandlerFunc {
	fd := EmbedFolder(m.fsData, m.target)
	indexHTML, indexErr := readEmbedFile(fd, indexFile)

	return func(c *gin.Context) {
		relativePath, valid := spaRelativePath(c.Request.URL.Path, urlPrefix)
		if !valid || strings.HasPrefix(relativePath, ".vite/") || strings.HasSuffix(relativePath, ".gz") {
			c.String(http.StatusNotFound, "404 page not found")
			c.Abort()
			return
		}

		if isIndexRequest(relativePath) {
			m.serveIndex(c, indexHTML, indexErr)
			return
		}

		// Runtime-generated files (for example, sitemap files) belong only to the primary SPA.
		if urlPrefix == "/" && m.serveStaticFile(relativePath, c) {
			c.Abort()
			return
		}

		if f, info, ok := openEmbedFile(fd, relativePath); ok {
			defer f.Close()
			if m.servePrecompressed(fd, relativePath, c) {
				c.Abort()
				return
			}
			serveContent(c, filepath.Base(relativePath), info.ModTime(), f)
			c.Abort()
			return
		}

		m.serveIndex(c, indexHTML, indexErr)
	}
}

func (m *SpaMdw) serveIndex(c *gin.Context, html []byte, indexErr error) {
	if indexErr != nil {
		c.String(http.StatusNotFound, "404 page not found")
		c.Abort()
		return
	}

	rtn := m.modifierIndex(html, c)
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
	c.Data(rtn.StatusCode, "text/html; charset=utf-8", rtn.Ret)
	c.Abort()
}

func (m *SpaMdw) servePrecompressed(fd ServeFileSystem, path string, c *gin.Context) bool {
	if !acceptsGzip(c.Request) || path == "" || path == indexFile {
		return false
	}

	f, info, ok := openEmbedFile(fd, path+".gz")
	if !ok {
		return false
	}
	defer f.Close()

	if contentType := mime.TypeByExtension(filepath.Ext(path)); contentType != "" {
		c.Header("Content-Type", contentType)
	}
	c.Header("Content-Encoding", "gzip")
	c.Writer.Header().Add("Vary", "Accept-Encoding")
	http.ServeContent(c.Writer, c.Request, filepath.Base(path), info.ModTime(), f)
	return true
}

func (m *SpaMdw) serveStaticFile(relativePath string, c *gin.Context) bool {
	if m.staticDir == "" {
		return false
	}

	root, err := os.OpenRoot(m.staticDir)
	if err != nil {
		return false
	}
	defer root.Close()

	f, err := root.Open(relativePath)
	if err != nil {
		return false
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil || info.IsDir() {
		return false
	}

	serveContent(c, filepath.Base(relativePath), info.ModTime(), f)
	return true
}

func acceptsGzip(r *http.Request) bool {
	var wildcardQuality *float64
	for _, value := range r.Header.Values("Accept-Encoding") {
		for _, encoding := range strings.Split(value, ",") {
			parts := strings.Split(encoding, ";")
			name := strings.TrimSpace(parts[0])
			if name == "" {
				continue
			}

			quality := 1.0
			for _, parameter := range parts[1:] {
				key, value, ok := strings.Cut(strings.TrimSpace(parameter), "=")
				if !ok || !strings.EqualFold(strings.TrimSpace(key), "q") {
					continue
				}
				parsed, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
				if err != nil || parsed < 0 || parsed > 1 {
					quality = 0
					break
				}
				quality = parsed
			}

			switch {
			case strings.EqualFold(name, "gzip"):
				return quality > 0
			case name == "*":
				wildcardQuality = &quality
			}
		}
	}
	return wildcardQuality != nil && *wildcardQuality > 0
}

func spaRelativePath(requestPath, urlPrefix string) (string, bool) {
	if urlPrefix != "" {
		if !strings.HasPrefix(requestPath, urlPrefix) {
			return "", false
		}
		requestPath = strings.TrimPrefix(requestPath, urlPrefix)
	}

	requestPath = strings.TrimPrefix(requestPath, "/")
	if requestPath == "" {
		return "", true
	}
	if !fs.ValidPath(requestPath) || path.Clean(requestPath) != requestPath {
		return "", false
	}
	return requestPath, true
}

func isIndexRequest(relativePath string) bool {
	return relativePath == "" || relativePath == indexFile
}

func serveContent(c *gin.Context, name string, modTime time.Time, content io.ReadSeeker) {
	if contentType := mime.TypeByExtension(filepath.Ext(name)); contentType != "" {
		c.Header("Content-Type", contentType)
	}
	http.ServeContent(c.Writer, c.Request, name, modTime, content)
}

func openEmbedFile(fileSystem ServeFileSystem, name string) (http.File, fs.FileInfo, bool) {
	f, err := fileSystem.Open(name)
	if err != nil {
		return nil, nil, false
	}

	info, err := f.Stat()
	if err != nil || info.IsDir() {
		_ = f.Close()
		return nil, nil, false
	}
	return f, info, true
}

func readEmbedFile(fileSystem ServeFileSystem, name string) ([]byte, error) {
	f, _, ok := openEmbedFile(fileSystem, name)
	if !ok {
		return nil, fs.ErrNotExist
	}
	defer f.Close()
	return io.ReadAll(f)
}

func (m *SpaMdw) modifierIndex(html []byte, c *gin.Context) *payloads.SeoPayload {
	rtn := &payloads.SeoPayload{Ret: html, StatusCode: http.StatusOK}
	if m.modifier == nil {
		return rtn
	}

	ret, err := m.modifier.SEO(html, c)
	if err != nil || ret == nil {
		return rtn
	}
	return ret
}

type embedFileSystem struct {
	http.FileSystem
}

func EmbedFolder(fileSystem fs.FS, targetPath string) ServeFileSystem {
	sub, err := fs.Sub(fileSystem, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{FileSystem: http.FS(sub)}
}

func Embed(fileSystem fs.FS) ServeFileSystem {
	return EmbedFolder(fileSystem, ".")
}
