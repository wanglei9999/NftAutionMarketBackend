package model

import (
	"time"
)

type SyncStatus struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null"` //名称
	LastSyncedBlock uint64 `gorm:"not null"` //最后同步的区块高度
	LastSyncedTime time.Time `gorm:"not null"` //最后同步时间
}