//建立路由结构
//为了便于测试几乎给每条路由都加了get方法
package Routers

import (
	"NetworkDisk/Handlers"
	"NetworkDisk/Middlewares"
	"NetworkDisk/Models/JWT"
	"NetworkDisk/Models/Redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildRouter(db *gorm.DB,pool *Redis.RedisPool,template JWT.Jwt) *gin.Engine {
	server:=gin.Default()

	group:=server.Group("/", Middlewares.CheakJWT(pool,template))
	{
		group.POST("/logout", Handlers.Logout(pool))

		group.POST("/upload", Handlers.Upload(db,pool))

		group.POST("/usesharedlinks", Handlers.Usesharedlinks(db,pool))

		group.GET("/getsharelinks", Handlers.Getsharelinks(pool))

		group.GET("/download", Handlers.Download(pool))

		group.GET("/filelist",Handlers.Filelist(pool))
	}

	server.POST("/register", Handlers.Register(db))

	server.POST("/login", Middlewares.CheakUserInfo(db),Handlers.Login(pool,template))

	return server
}
