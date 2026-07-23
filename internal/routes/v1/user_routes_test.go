package v1_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/Xwudao/neter-template/internal/biz/mocks"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/data/ent"
	v1 "github.com/Xwudao/neter-template/internal/routes/v1"
)

func init() {
	gin.SetMode(gin.TestMode)
}

// newTestUserRoute builds a UserRoute with a mock biz and returns the gin engine.
func newTestUserRoute(t *testing.T, mockBiz *mocks.MockUserBizIface) *gin.Engine {
	t.Helper()
	g := gin.New()
	conf := koanf.New(".")
	route := v1.NewUserRoute(g, mockBiz, conf)
	route.Register()
	return g
}

// TestUserRoute_Login_Success verifies the login handler returns a token on valid credentials.
func TestUserRoute_Login_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockBiz := mocks.NewMockUserBizIface(ctrl)
	mockBiz.EXPECT().
		Login(gomock.Any(), gomock.Any()).
		Return(&ent.User{ID: 1}, "test-token", nil)

	g := newTestUserRoute(t, mockBiz)

	w := httptest.NewRecorder()
	body := `{"username":"admin","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/v1/user/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp core.WrappedResp
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, core.CodeSuccess, resp.Code)
}

// TestUserRoute_Login_InvalidJSON verifies the handler rejects malformed JSON bodies.
func TestUserRoute_Login_InvalidJSON(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Login should never be called when binding fails.
	mockBiz := mocks.NewMockUserBizIface(ctrl)
	mockBiz.EXPECT().Login(gomock.Any(), gomock.Any()).Times(0)

	g := newTestUserRoute(t, mockBiz)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/user/login", strings.NewReader("{bad json"))
	req.Header.Set("Content-Type", "application/json")
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp core.WrappedResp
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	// Handler wraps binding errors as a non-success code.
	assert.NotEqual(t, core.CodeSuccess, resp.Code)
}

// TestUserRoute_Login_BizError verifies the handler propagates biz-layer errors.
func TestUserRoute_Login_BizError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockBiz := mocks.NewMockUserBizIface(ctrl)
	mockBiz.EXPECT().
		Login(gomock.Any(), gomock.Any()).
		Return(nil, "", assert.AnError)

	g := newTestUserRoute(t, mockBiz)

	w := httptest.NewRecorder()
	body := `{"username":"admin","password":"wrong"}`
	req := httptest.NewRequest(http.MethodPost, "/v1/user/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp core.WrappedResp
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.NotEqual(t, core.CodeSuccess, resp.Code)
}
