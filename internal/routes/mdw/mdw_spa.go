package mdw

import (
	"embed"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/domain/payloads"
	"github.com/Xwudao/neter-template/pkg/libx"
)

const INDEX = "index.html"

type ServeFileSystem interface {
	http.FileSystem
	Exists(prefix string, path string) bool
}

type SpaMdw struct {
	fsData embed.FS
	target string

	sb *biz.SeoBizBiz
	mf libx.ViteManifest
}

func NewSpaMdw(fsData embed.FS, target string, sb *biz.SeoBizBiz) (*SpaMdw, error) {
	sm := &SpaMdw{
		fsData: fsData,
		target: target,
		sb:     sb,
	}

	err := sm.parseManifest()
	if err != nil {
		return nil, err
	}

	return sm, nil
}

// parseManifest 解析manifest.json
func (m *SpaMdw) parseManifest() error {
	fd := EmbedFolder(m.fsData, m.target)
	mFile, err := fd.Open(".vite/manifest.json")
	if err != nil {
		return err
	}
	defer mFile.Close()
	mCnt, _ := io.ReadAll(mFile)

	manifestString, err := libx.ParseManifestString(string(mCnt))
	if err != nil {
		return err
	}

	m.mf = manifestString
	return nil
}

// acceptsGzip 检查客户端是否支持 gzip 压缩
func (m *SpaMdw) acceptsGzip(c *gin.Context) bool {
	acceptEncoding := c.GetHeader("Accept-Encoding")
	return strings.Contains(strings.ToLower(acceptEncoding), "gzip")
}

// servePrecompressed 尝试提供预压缩的静态文件
func (m *SpaMdw) servePrecompressed(c *gin.Context, path string) bool {
	// 检查客户端是否支持 gzip
	if !m.acceptsGzip(c) {
		return false
	}

	fd := EmbedFolder(m.fsData, m.target)
	gzPath := path + ".gz"

	// 检查是否存在预压缩文件
	if !fd.Exists("", gzPath) {
		return false
	}

	// 打开并读取预压缩文件
	gzFile, err := fd.Open(gzPath)
	if err != nil {
		return false
	}
	defer gzFile.Close()

	gzContent, err := io.ReadAll(gzFile)
	if err != nil {
		return false
	}

	// 设置正确的响应头
	mimeType := mime.TypeByExtension(filepath.Ext(path))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	// 设置 gzip 相关的响应头
	c.Header("Content-Encoding", "gzip")
	c.Header("Content-Type", mimeType)
	c.Header("Vary", "Accept-Encoding")

	// 可选：设置缓存头
	c.Header("Cache-Control", "public, max-age=604800") // 1周缓存

	c.Data(http.StatusOK, mimeType, gzContent)
	return true
}

func (m *SpaMdw) Serve(urlPrefix string) gin.HandlerFunc {
	fd := EmbedFolder(m.fsData, m.target)
	fileServer := http.FileServer(fd)
	if urlPrefix != "" {
		fileServer = http.StripPrefix(urlPrefix, fileServer)
	}

	return func(c *gin.Context) {
		var path = c.Request.URL.Path
		switch {
		case m.mf.IsStaticFileExists(path):
			// 尝试提供 gzip 预压缩文件
			if m.servePrecompressed(c, path) {
				c.Abort()
				return
			}
			// 如果没有预压缩文件或客户端不支持压缩，使用原始文件
			fileServer.ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return

		case m.hasStaticFile(path):
			cnt, mimeType := m.readStaticFile(path, c)
			c.Data(http.StatusOK, mimeType, cnt)
			c.Abort()
			return

		default:
			rtn := m.modifierIndex(fd, c)
			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Header("Pragma", "no-cache")
			c.Header("Expires", "0")
			c.Data(rtn.StatusCode, "text/html; charset=utf-8", rtn.Ret)
			c.Abort()
			return
		}
	}
}

func (m *SpaMdw) hasStaticFile(path string) bool {
	pw, _ := os.Getwd()
	// 检查 gzip 文件
	gzPath := filepath.Join(pw, "static", path+".gz")
	if info, err := os.Stat(gzPath); err == nil && !info.IsDir() {
		return true
	}
	// 检查原始文件
	fp := filepath.Join(pw, "static", path)
	info, err := os.Stat(fp)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func (m *SpaMdw) readStaticFile(path string, c *gin.Context) ([]byte, string) {
	pw, _ := os.Getwd()

	// 优先尝试 gzip 文件
	if m.acceptsGzip(c) {
		gzPath := filepath.Join(pw, "static", path+".gz")
		if gzCnt, err := os.ReadFile(gzPath); err == nil {
			mimeType := mime.TypeByExtension(filepath.Ext(path))
			// 设置 gzip 相关的响应头
			c.Header("Content-Encoding", "gzip")
			c.Header("Vary", "Accept-Encoding")
			c.Header("Cache-Control", "public, max-age=604800")
			return gzCnt, mimeType
		}
	}

	// 最后使用原始文件
	fp := filepath.Join(pw, "static", path)
	cnt, err := os.ReadFile(fp)
	if err != nil {
		return nil, ""
	}
	mimeType := mime.TypeByExtension(filepath.Ext(fp))

	return cnt, mimeType
}

func (m *SpaMdw) modifierIndex(fs ServeFileSystem, c *gin.Context) *payloads.SeoPayload {
	var rtn = &payloads.SeoPayload{
		StatusCode: http.StatusOK,
		Ret:        nil,
	}

	f, err := fs.Open(INDEX)
	if err != nil {
		return rtn
	}
	defer f.Close()

	all, err := io.ReadAll(f)
	if err != nil {
		return rtn
	}

	ret, _ := m.sb.SEO(all, c)
	if ret == nil {
		return rtn
	}
	return ret

}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) ServeFileSystem {
	sub, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(sub),
	}
}

func Embed(fsEmbed embed.FS) ServeFileSystem {
	return EmbedFolder(fsEmbed, ".")
}
