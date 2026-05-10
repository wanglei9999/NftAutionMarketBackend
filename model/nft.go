package model

import (
	"time"
	"gorm.io/gorm"
)

type NFT struct {
	gorm.Model
	ContractAddress string `gorm:"type:varchar(42);not null"` //NFT合约地址
	TokenID string `gorm:"type:varchar(100);not null"` //NFT的唯一标识符
	MetadataURI string `gorm:"type:varchar(255);not null"` //NFT的元数据URI
	OwnerAddress string `gorm:"type:varchar(42);not null"` //当前拥有者地址
	Name string `gorm:"type:varchar(100);not null"` //NFT名称
	Description string `gorm:"type:varchar(255);not null"` //NFT描述
	ImageURI string `gorm:"type:varchar(255);not null"` //NFT图片URI
	ChainID uint64 `gorm:"not null"` //链ID
}