package main

import (
	"NetworkDisk/Models/JWT"
	"NetworkDisk/config"
)

const ConfPath 	= "conf.json"	//外部配置文件
const Secret 	= "I'mTooDishes"//JWT秘钥
var conf config.Conf

var (
	DefaultJwt = JWT.Jwt{
		Header: JWT.Header{
			Alg: "HS256",
			Typ: "JWT",
		},
		Payload: JWT.Payload{
			Iss: "kip",
			Sub: "NetworkDisk",
		},
		Secret: Secret,
	}
)