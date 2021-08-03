package File

import "gorm.io/gorm"

type Privilege struct {
	Pri 	 uint 	`gorm:"primaryKey"`
	PathName string `gorm:"not null"`
	Owner    int    `gorm:"not null"` //拥有权用户ID
	User	 int//有使用权的用户
}

func (p *Privilege)Save(db *gorm.DB) error {
	var tmp Privilege
	db=db.Where(&Privilege{PathName: p.PathName,User: p.User}).Find(&tmp)
	if db.Error!=nil {
		return db.Error
	}
	if tmp.Pri!=0 {//同一条记录
		return nil
	}
	db=db.Save(p)
	return db.Error
}