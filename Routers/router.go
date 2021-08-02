//建立路由结构
//为了便于测试几乎给每条路由都加了get方法
package Routers

import (
	"NetworkDisk/Handlers"
	"NetworkDisk/Middlewares"
	"NetworkDisk/Models/jwt"
	"NetworkDisk/Models/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func BuildRouter(db *gorm.DB,pool redis.RedisPool,template jwt.Jwt) *gin.Engine {
	server:=gin.Default()

	group:=server.Group("/", Middlewares.CheakJWT(pool,template))
	{
		group.POST("/logout", Handlers.Logout(pool))
		group.GET("/logout",func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "logout",
			})
		})

		group.POST("/upload", Handlers.Upload())
		group.GET("/upload",func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "logout",
			})
		})

		group.GET("/download", Handlers.Download())
		//group.GET("/download",func(c *gin.Context) {
		//	c.JSON(http.StatusOK,gin.H{
		//		"method":  "GET",
		//		"routing": "download",
		//	})
		//})
	}

	server.POST("/test",Handlers.Test(pool))

	server.POST("/register", Handlers.Register(db))
	server.GET("/register",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":  "GET",
			"routing": "register",
		})
	})

	server.POST("/login", Middlewares.CheakUserInfo(db),Handlers.Login(pool,template))
	server.GET("/login",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":  "GET",
			"routing": "login",
		})
	})

	return server
}
