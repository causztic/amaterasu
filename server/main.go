package main

import (
	"time"

	"./models"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

var authMiddleware *jwt.GinJWTMiddleware

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"email": claims["email"],
		"role":  claims["role"],
		"text":  "Hello World.",
	})
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
		v1.Use(authMiddleware.MiddlewareFunc())
		{
			v1.GET("/hello", helloHandler)
		}
	}

	return r
}

func setupAuth() {
	// the jwt middleware
	authMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(email string, password string, c *gin.Context) (string, bool) {
			if models.AuthenticateCredentials(email, password) {
				return email, true
			}
			return email, false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}

func main() {
	models.InitDB()
	setupAuth()
	r := setupRouter()
	r.Run(":9000")
}
