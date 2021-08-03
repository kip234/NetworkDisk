package Handlers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

//用于从上下文中取出UID
func getUid(c *gin.Context) (uid int,err error) {
	v,ok:=c.Get("uid")
	if !ok {
		uid=-1
		err = fmt.Errorf("Missing UID field")
		return
	}
	uid,ok = v.(int)
	if !ok {
		uid=-1
		err = fmt.Errorf("Assertion failure")
		return
	}
	return uid,nil
}

//获取分享链接
func Encoding(owner int,u int,path,name string) string {
	s:=strconv.Itoa(owner)+"*"+path+"*"+strconv.Itoa(u)+"*"+name
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Decoding(link string) (owner int,u int,path,name string,err error) {
	v,err:=base64.StdEncoding.DecodeString(link)//解码
	if err!=nil {
		return 0,0,"","",err
	}
	s:=string(v)
	data:=strings.Split(s,"*")//分割信息
	if len(data)!=4 {
		return 0,0,"","",fmt.Errorf("link error")
	}
	owner,err = strconv.Atoi(data[0])//解析所有者
	if err!=nil {
		return 0,0,"","",err
	}
	u,err = strconv.Atoi(data[2])//解析目标用户
	if err!=nil {
		return 0,0,"","",err
	}
	path=data[1]
	name=data[3]
	return
}