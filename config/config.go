package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Sql struct {
	SqlName 	string	//数据库名
	SqlUserName string	//数据库登录用账户名
	SqlUserPwd 	string	//数据库登录用账户密码
	SqlAddr 	string	//数据库地址
}

type Conf struct {
	Sql
	Rredis 		string	//Redis地址
	Wredis		string	//Redis地址
	Addr string	//服务器地址
}

func Init(ConfPath string) (result Conf) {
	file,err:=os.Open(ConfPath)
	if err!=nil {
		panic(err)
	}
	buf,err:=ioutil.ReadAll(file)
	if err!=nil {
		panic(err)
	}
	json.Unmarshal(buf,&result)
	return
}