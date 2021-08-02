package Handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
