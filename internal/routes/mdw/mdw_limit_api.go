package mdw

import (
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

type ApiLimitOpt struct {
	Path            string `json:"path"`
	CountPerSeconds int    `json:"count_per_seconds"`
	Slack           int    `json:"slack"`
}

func ApiLimitMiddleware(opts []ApiLimitOpt) gin.HandlerFunc {
	var syncMap sync.Map
	for _, opt := range opts {
		syncMap.LoadOrStore(opt.Path, ratelimit.New(opt.CountPerSeconds, ratelimit.WithSlack(opt.Slack)))
	}
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/assets/") {
			return
		}
		pt := c.Request.URL.Path
		limiter, ok := syncMap.Load(pt)
		if !ok {
			c.Next()
			return
		}
		limiter.(ratelimit.Limiter).Take()
		c.Next()
	}
}
