package mdw

import (
	"bytes"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DumpReqResMdw(enable bool, log *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !enable {
			return
		}
		var (
			path       = c.Request.URL.Path
			query      = c.Request.URL.Query().Encode()
			method     = c.Request.Method
			name       = c.GetHeader("x-name")
			reqBodyStr string
			resBodyStr string

			nowtime = time.Now().Format("20060102150405")
		)

		if name == "" {
			return
		}

		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		reqBodyStr = string(bodyBytes)

		c.Writer = &bodyWriter{bodyCache: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Next()

		bw, ok := c.Writer.(*bodyWriter)
		if ok {
			resBodyStr = bw.bodyCache.String()
		}

		filename := nowtime + "_" + strings.ToLower(name) + "_" + strings.ReplaceAll(path, "/", "_") + ".log"
		aimFolder := "./build"
		if _, err := os.Stat(aimFolder); err != nil {
			_ = os.Mkdir(aimFolder, os.ModePerm)
		}

		f, err := os.OpenFile(aimFolder+"/"+filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}

		defer f.Close()

		var sb = strings.Builder{}
		sb.WriteString("path: " + path + "\n")
		sb.WriteString("query: " + query + "\n")
		sb.WriteString("name: " + name + "\n")
		sb.WriteString("method: " + method + "\n")
		sb.WriteString("reqbody: " + strings.ReplaceAll(reqBodyStr, "\n", "") + "\n")
		sb.WriteString("resbody: " + resBodyStr + "\n")

		_, _ = f.WriteString(sb.String())
		log.Infof("dump request and response to %s", filename)
	}
}

type bodyWriter struct {
	gin.ResponseWriter
	bodyCache *bytes.Buffer
}

// rewrite Write()
func (w bodyWriter) Write(b []byte) (int, error) {
	w.bodyCache.Write(b)
	return w.ResponseWriter.Write(b)
}
