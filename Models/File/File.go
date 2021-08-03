package File

import (
	"gorm.io/gorm"
)

type File struct {
	PathName string `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Size     int64   `gorm:"not null"`
	Owner    int    `gorm:"not null"` //拥有权用户ID
}

func (f *File)Save(db *gorm.DB) error {
	db = db.Save(f)//刷新
	return db.Error
}