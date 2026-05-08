package model

import(
	"time"
	"gorm.io/gorm"
)

type Transaction struct{
	gorm.Model
	TxHash string `gorm:"type:char(66);uniqueIndex;not null"`
	From string `gorm:"type:char(42);index"`
	To string `gorm:"type:char(42);index"`
	Method string `gorm:"type:varchar(64)"`
	BlockNum unit64 `gorm:"index"`
	LogIndex unit `gorm:"index"`
	ChainID unit64 `gorm:"index"`
	Timestamp time.Time
}