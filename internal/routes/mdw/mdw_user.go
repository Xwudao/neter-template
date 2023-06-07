package mdw

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/data"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/data/ent/user"
	"github.com/Xwudao/neter-template/pkg/enum"
	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

// ExtractUserInfoMiddleware just extract the user info from the request, and save it to the context.
func ExtractUserInfoMiddleware(logger *zap.SugaredLogger, jc *jwt.Client, data *data.Data) gin.HandlerFunc {
	return func(c *gin.Context) {
		var logged bool
		defer func() {
			c.Header("X-Logged", strconv.FormatBool(logged))
		}()
		authHeader := c.GetHeader(enum.KeyAuthorization)
		if authHeader == "" {
			c.Next()
			return
		}

		claims, err := jc.Parse(strings.ReplaceAll(authHeader, enum.KeyBearer, ""))
		if err != nil {
			logger.Errorf("parse token failed: %v", err)
			c.Next()
			return
		}

		logger.Infof("user cliams: %#v", claims)

		userID := cast.ToInt64(claims["user_id"])
		logger.Infof("user id: %d", userID)
		//userID = int64(claims["user_id"].(float64))
		client := data.Client

		userInfo, err := client.User.Query().Where(user.ID(userID)).First(c.Request.Context())
		if err != nil {
			c.Next()
			return
		}
		logged = true
		logger.Infof("user [%s] logged", userInfo.Username)
		c.Set(enum.KeyUserInfo, userInfo)
	}
}

//func ExtractUserInfoByTokenMdw(data *data.Data) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var logged bool
//		defer func() {
//			c.Header("X-Logged", strconv.FormatBool(logged))
//		}()
//
//		token := c.Query("token")
//		if token == "" {
//			c.Next()
//			return
//		}
//
//		us, err := data.Client.User.Query().Where(user.Token(token)).First(c.Request.Context())
//		if err != nil {
//			c.Next()
//			return
//		}
//
//		logged = true
//		c.Set(enum.KeyUserInfo, us)
//	}
//}

func MustLoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get(enum.KeyUserInfo)
		if !exists {
			//c.AbortWithStatus(http.StatusForbidden)
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"code": 403,
					"msg":  "please login first",
					"data": nil,
				},
			)
			return
		}
	}
}

func MustWithRoleMiddleware(needRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, exists := c.Get(enum.KeyUserInfo)
		if !exists {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"code": 403,
					"msg":  "please login first",
					"data": nil,
				},
			)
			return
		}

		userInfo := u.(*ent.User)
		if userInfo.Role != needRole {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"code": 403,
					"msg":  "you don't have permission to do this",
					"data": nil,
				},
			)
			return
		}

		c.Next()
		return

	}
}
