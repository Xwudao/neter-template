package mdw

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/data"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/data/ent/user"
	"github.com/Xwudao/neter-template/pkg/enum"
	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

//ExtractUserInfoMiddleware just extract the user info from the request, and save it to the context.
func ExtractUserInfoMiddleware(logger *zap.SugaredLogger, jc *jwt.Client, data *data.Data) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(enum.KeyAuthorization)
		if authHeader == "" {
			c.Next()
			return
		}

		claims, err := jc.Parse(strings.ReplaceAll(authHeader, enum.KeyBearer, ""))
		if err != nil {
			c.Next()
			return
		}

		userID := int64(claims["user_id"].(float64))
		client := data.Client

		userInfo, err := client.User.Query().Where(user.ID(userID)).First(c.Request.Context())
		if err != nil {
			c.Next()
			return
		}
		logger.Infof("user [%s] login", userInfo.Username)
		c.Set(enum.KeyUserInfo, userInfo)
	}
}

func MustLoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get(enum.KeyUserInfo)
		if !exists {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}

func MustWithRoleMiddleware(needRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, exists := c.Get(enum.KeyUserInfo)
		if !exists {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		userInfo := u.(*ent.User)
		if userInfo.Role != needRole {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
		return

	}
}
