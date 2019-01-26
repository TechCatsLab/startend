package middlewares

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/TechCatsLab/startend/backend/services"
)

type WxSession struct {
	SessionKey string `json:"session_key"`
	ExpireIn   int    `json:"expires_in"`
	OpenID     string `json:"openid"`
}

var (
	// JWTMiddleware should be exported for user authentication.
	JWTMiddleware *jwt.GinJWTMiddleware

	// skipMap = map[string]bool{
	// 	"/api/v1/user/login": true,
	// }
)

func init() {
	JWTMiddleware = &jwt.GinJWTMiddleware{
		Realm:   "Template",
		Key:     []byte("task"),
		Timeout: 24 * time.Hour,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var req struct {
				Jscode string `json:"jscode"`
			}

			if err := c.BindJSON(&req); err != nil {
				glog.Errorf("[status] parameter error:%s", err)

				return nil, err
			}
			client := &http.Client{}
			url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", "wxd37a86f36f703828", "980cc565babd9ba31e76c1a597270117", req.Jscode)
			reqest, err := http.NewRequest("GET", url, nil)
			if err != nil {

				return nil, err
			}

			response, err := client.Do(reqest)
			if err != nil {
				glog.Errorf("[status] WxRequest error:%s", err)

				return nil, err
			}
			var session WxSession
			body, err := ioutil.ReadAll(response.Body)
			if err := json.Unmarshal(body, &session); err != nil {

				return nil, err
			}

			uid, err := services.UserService.QueryUid(session.OpenID)
			if err != nil {
				glog.Errorf("[status] uid get error:%s", err)

				return nil, jwt.ErrFailedAuthentication
			}

			return int(uid), nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(int); ok {

				return true
			}

			return false
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int); ok {
				return jwt.MapClaims{
					"id": v,
				}
			}

			return jwt.MapClaims{}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(403, gin.H{
				"code":    403,
				"message": "You don't have permission to access.",
			})
		},
	}
}

// func installJWTMiddleware(router *gin.Engine) {
// 	router.Use(func(c *gin.Context) {
// 		glog.Infof("[JWT] Middleware, path %s", c.Request.URL.Path)
// 		if _, skip := skipMap[c.Request.URL.Path]; skip {
// 			c.Next()
// 			return
// 		}

// 		JWTMiddleware.MiddlewareFunc()(c)
// 	})
// }
