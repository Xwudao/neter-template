package mdw

import (
	"bytes"
	"embed"
	"io"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

const INDEX = "index.html"

type ServeFileSystem interface {
	http.FileSystem
	Exists(prefix string, path string) bool
}

type SpaMdw struct {
	fsData embed.FS
	target string
}

func NewSpaMdw(fsData embed.FS, target string) *SpaMdw {
	return &SpaMdw{fsData: fsData, target: target}
}

func (m *SpaMdw) Serve(urlPrefix string) gin.HandlerFunc {
	fd := EmbedFolder(m.fsData, m.target)
	fileServer := http.FileServer(fd)
	if urlPrefix != "" {
		fileServer = http.StripPrefix(urlPrefix, fileServer)
	}
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/" {
			c.Data(200, "text/html; charset=utf-8", m.modifierIndex(fd))
			c.Abort()
			return
		}
		if fd.Exists(urlPrefix, c.Request.URL.Path) {
			fileServer.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

func (m *SpaMdw) ServeNotFound(fallbackFile string) gin.HandlerFunc {
	if fallbackFile == "" {
		fallbackFile = INDEX
	}
	fd := EmbedFolder(m.fsData, m.target)
	return func(c *gin.Context) {
		f, err := fd.Open(fallbackFile)
		if err != nil {
			c.String(http.StatusNotFound, "404 page not found")
			return
		}
		defer f.Close()
		c.Data(http.StatusOK, "text/html; charset=utf-8", m.modifierIndex(fd))
	}
}

func (m *SpaMdw) modifierIndex(fs ServeFileSystem) []byte {
	f, err := fs.Open(INDEX)
	if err != nil {
		return nil
	}
	defer f.Close()

	all, err := io.ReadAll(f)
	if err != nil {
		return nil
	}

	// you can replace something here
	return bytes.ReplaceAll(all, []byte(""), []byte(""))
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
