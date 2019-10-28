package middlewares

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/bindu-bindu/bindu-blank/app/models"
	"github.com/bindu-bindu/bindu-blank/db"
	"github.com/gin-gonic/gin"
)

// GinJwtMiddlewareHandler to handle authentication and authorization for api routes
func GinJwtMiddlewareHandler() *jwt.GinJWTMiddleware {
	var identityKey = "id"
	// the jwt middleware

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Username: fmt.Sprintf("%v", claims[identityKey]),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// Only login will fire here
			var loginVals models.User
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			var db = db.DB

			if err := db.Where("username=?", userID).Find(&loginVals).Limit(1).Error; err != nil {
				c.AbortWithStatus(401)
				fmt.Println(err)
			} else {
				// fmt.Println(loginVals)
				if userID == loginVals.Username && password == loginVals.Password {
					return &loginVals, nil
				}
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// if v, ok := data.(*models.User); ok && v.Username == "admin" {
			var user models.User
			v, ok := data.(*models.User)
			if ok {
				var db = db.DB
				if err := db.Where("username=?", v.Username).Find(&user).Limit(1).Error; err != nil {
					c.AbortWithStatus(401)
					fmt.Println(err)
					return false
				}
				// adapter, _ := gormadapter.NewAdapterByDB(db)

				// e, _ := casbin.NewEnforcer("middlewares/auth_model.conf", adapter)

				// Load policies from DB dynamically
				// _ = e.LoadPolicy()
				// fmt.Println(e)
				// if err != nil {
				// 	return false, fmt.Errorf("failed to create casbin enforcer: %w", e)
				// }

				// ok := e.Enforce(sub, obj, act)
				// fmt.Println(ok)
				// return ok, err

				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"Code":    code,
				"Message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}
