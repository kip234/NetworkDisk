package main

import (
	"NetworkDisk/Database"
	"NetworkDisk/Models/File"
	"NetworkDisk/Models/Redis"
	"NetworkDisk/Routers"
	"NetworkDisk/config"
)

func main()  {
	//s:=gin.Default()
	//s.Run("localhost:8080")

	conf=config.Init(ConfPath)//获取服务器配置
	var redis = Redis.RedisPool{
		Write:		conf.Wredis,
		Read:		conf.Rredis,
		IdLeTimeout:5,
		MaxIdle:	20,
		MaxActive:	8,
	}
	redis.Init()
	//fmt.Println(conf)
	db:=Database.InitGorm(&conf.Sql)
	//database.
	//缓存数据
	err := File.Out(db,redis)
	if err!=nil{
		panic(err)
	}
	server:=Routers.BuildRouter(db,&redis,DefaultJwt)
	server.Run(conf.Addr)
}
