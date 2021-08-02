package Models

import (
	"gorm.io/gorm"
)

type User struct {
	Uid int `gorm:"primaryKey"`
	Name string	`gorm:"string not null"`//用户名
	Pwd string `gorm:"string not null"`//用户密码
}

//创建用户
func (u *User)Save(db *gorm.DB) (err error) {
	err=db.Create(u).Error
	return
}

//根据提供的UID读取用户信息
//主要用于密码比对
func (u *User)Load(db *gorm.DB,uid int) (err error) {
	err=db.Where("Uid=?",uid).Find(u).Error
	return
}

//判断密码是否正确，如果不正确返回false
func (u *User)PwdIsRight(db *gorm.DB) bool {
	tmp:=User{}
	db.Where("Uid=?", u.Uid).Find(&tmp)
	return tmp.Pwd==u.Pwd
}

//判断是否存在，如果不存在返回false
func (u *User)IsExist(db *gorm.DB) bool {
	tmp:=User{}
	db.Where("Uid=?", u.Uid).Find(&tmp)
	if nil == db.Error {
			return true
		}
	return false
}
