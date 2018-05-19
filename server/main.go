package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"

	"./fs"
	"./models"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/rtmp"
)

var authMiddleware *jwt.GinJWTMiddleware
var mutex *sync.RWMutex
var rtmpServer *rtmp.Server

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"email": claims["email"],
		"role":  claims["role"],
	})
}

func itemsHandler(c *gin.Context) {
	dir := c.DefaultQuery("dir", "/")

	if len(dir) == 0 {
		dir = "/"
	}

	items := fs.GetDirectoryItems(dir)
	c.JSON(200, items)
}

func setupRmtp() {
	rtmpServer = &rtmp.Server{}
	mutex = &sync.RWMutex{}
	type Channel struct {
		que *pubsub.Queue
	}
	channels := map[string]*Channel{}

	rtmpServer.HandlePlay = func(conn *rtmp.Conn) {
		mutex.RLock()
		ch := channels[conn.URL.Path]
		mutex.RUnlock()

		if ch != nil {
			cursor := ch.que.Latest()
			avutil.CopyFile(conn, cursor)
		}
	}

	rtmpServer.HandlePublish = func(conn *rtmp.Conn) {
		streams, _ := conn.Streams()

		mutex.Lock()
		ch := channels[conn.URL.Path]
		if ch == nil {
			ch = &Channel{}
			ch.que = pubsub.NewQueue()
			ch.que.WriteHeader(streams)
			channels[conn.URL.Path] = ch
		} else {
			ch = nil
		}
		mutex.Unlock()
		if ch == nil {
			return
		}

		avutil.CopyPackets(ch.que, conn)

		mutex.Lock()
		delete(channels, conn.URL.Path)
		mutex.Unlock()
		ch.que.Close()
	}

}

func itemHandler(c *gin.Context) {
	name := c.Query("name")
	split := strings.Split(name, ".")
	split2 := strings.Split(name, "/")
	lastName := split2[len(split2)-1]
	extension := split[len(split)-1]
	if extension == "avi" {
		// run ffmpeg to start
		cmd := exec.Cmd{
			Path: "ffmpeg",
			Args: []string{"-re", "-i", name, "-c copy", "-f", "flv", "rtmp://localhost/", lastName},
		}
		go func() {
			_ = cmd.Run()
		}()
		c.JSON(200, gin.H{
			"url": fmt.Sprintf("rtmp://localhost/%s", lastName),
		})
	} else {
		c.File(name)
	}
	// c.JSON(200, gin.H{
	// 	"message": fmt.Sprintf("Not a MP4: is a %s", extension),
	// })
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authMiddleware.LoginHandler)
		}
		v1.GET("/item", itemHandler)
		v1.Use(authMiddleware.MiddlewareFunc())
		{
			v1.GET("/hello", helloHandler)
			v1.GET("/items", itemsHandler)
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
		Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
			if models.AuthenticateCredentials(username, []byte(password)) {
				return username, true
			}
			return username, false
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
	setupRmtp()
	go rtmpServer.ListenAndServe()
	r := setupRouter()
	r.Run(":9000")
}
