package mdw

import (
	"embed"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"go-kitboxpro/internal/biz"
	"go-kitboxpro/internal/domain/payloads"
	"go-kitboxpro/pkg/libx"
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

func (m *SpaMdw) Serve(urlPrefix string) gin.HandlerFunc {
	fd := EmbedFolder(m.fsData, m.target)
	fileServer := http.FileServer(fd)
	if urlPrefix != "" {
		fileServer = http.StripPrefix(urlPrefix, fileServer)
	}

	return func(c *gin.Context) {
		var path = c.Request.URL.Path
		switch {
		//case fd.Exists(urlPrefix, c.Request.URL.Path) && path != "/":
		case m.mf.IsStaticFileExists(path):
			fileServer.ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return

		case m.hasStaticFile(path):
			cnt, mimeType := m.readStaticFile(path)
			c.Data(http.StatusOK, mimeType, cnt)

		default:
			//if !fd.Exists("", INDEX) {
			//	c.String(http.StatusNotFound, "404 page not found")
			//	return
			//}
			rtn := m.modifierIndex(fd, c)
			c.Data(rtn.StatusCode, "text/html; charset=utf-8", rtn.Ret)
			c.Abort()
			return
		}

	}
}

func (m *SpaMdw) hasStaticFile(path string) bool {
	pw, _ := os.Getwd()
	fp := filepath.Join(pw, "static", path)
	info, err := os.Stat(fp)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func (m *SpaMdw) readStaticFile(path string) ([]byte, string) {
	pw, _ := os.Getwd()
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
