package mdw

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"sync"
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
