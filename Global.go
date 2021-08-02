package main

import (
	"NetworkDisk/Models/jwt"
	"NetworkDisk/Models/redis"
)

const ConfPath 	= "conf.json"	//外部配置文件
const Secret 	= "I'mTooDishes"//JWT秘钥
const (
	Rredis = "localhost:6379"
	Wredis = "localhost:6379"
)

var (
	DefaultJwt = jwt.Jwt{
		Header: jwt.Header{
			Alg: "HS256",
			Typ: "JWT",
		},
		Payload: jwt.Payload{
			Iss: "kip",
			Sub: "NetworkDisk",
		},
		Secret: Secret,
	}
)

var Redis = redis.RedisPool{
	Write:		Rredis,
	Read:		Rredis,
	IdLeTimeout:5,
	MaxIdle:	20,
	MaxActive:	8,
}