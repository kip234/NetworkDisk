package main

import (
	"NetworkDisk/Database"
	"NetworkDisk/Routers"
	"NetworkDisk/config"
)

func main()  {
	//s:=gin.Default()
	//s.Run("localhost:8080")
	Redis.Init()
	conf:=config.Init(ConfPath)//获取服务器配置
	db:=Database.InitGorm(&conf.Sql)
	//database.
	server:=Routers.BuildRouter(db,Redis,DefaultJwt)
	server.Run(conf.Addr)
}
