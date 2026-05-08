package model

import (
	"gorm.io/gorm"
)


//用户
type User struct{
	gorm.model
	Address string `gorm:"type:char(42);uniqueIndex;not null"`
	NickName string `gorm:"type"varchar(255)"`
}