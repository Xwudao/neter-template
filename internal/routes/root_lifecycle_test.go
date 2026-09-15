package routes

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Xwudao/loom"
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/system"
)

func TestHTTPServerLifecycle(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/health", func(c *gin.Context) { c.Status(http.StatusNoContent) })
	conf := koanf.New(".")
	require.NoError(t, conf.Set("app.port", 0))
	require.NoError(t, conf.Set("app.host", true))
	appCtx := system.NewTestAppContext()
	lifecycle := loom.NewLifecycle()
	engine, err := NewHttpEngine(router, conf, zap.NewNop().Sugar(), appCtx, nil, nil, lifecycle)
	require.NoError(t, err)

	require.NoError(t, lifecycle.Start(t.Context()))
	resp, err := http.Get("http://" + engine.server.Addr + "/health")
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
	require.NoError(t, resp.Body.Close())

	shutdownCtx, cancel := context.WithTimeout(t.Context(), time.Second)
	defer cancel()
	require.NoError(t, lifecycle.Stop(shutdownCtx))
	select {
	case <-appCtx.Ctx.Done():
	case <-time.After(time.Second):
		t.Fatal("application context was not cancelled")
	}
	_, err = http.Get("http://" + engine.server.Addr + "/health")
	require.Error(t, err)
}
