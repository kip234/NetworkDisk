package File

import (
	"NetworkDisk/Models/Redis"
	"gorm.io/gorm"
	"strconv"
)

//得到使用者Redis的键
func UserKey(id int) string {
	return "User"+strconv.Itoa(id)
}

//得到拥有者Redis的键
func OwnerKey(id int) string {
	return "Owner"+strconv.Itoa(id)
}

//输出所有记录到Redis
func Out(db *gorm.DB,redis Redis.RedisPool) (err error) {
	var p []Privilege
	db=db.Find(&p)
	if db.Error!=nil {
		return db.Error
	}
	for _,i:=range p{
		err=redis.SADD(UserKey(i.User),i.PathName) //记录可以使用的
		if err!=nil {
			return err
		}
		err=redis.SADD(OwnerKey(i.Owner),i.PathName) //记录拥有的
		if err!=nil {
			return err
		}
	}
	return nil
}