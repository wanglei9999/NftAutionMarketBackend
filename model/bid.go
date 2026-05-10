package model

import (
	"time"
	"gorm.io/gorm"
)

type Bid struct {
	gorm.Model
	AuctionID uint `gorm:"not null"` //外键关联到Auction表
	BidderAddress string `gorm:"type:varchar(42);not null"` //出价者地址
	TokenAddress string `gorm:"type:varchar(42);not null"` //ETH 或ERC20代币地址
	Amount string `gorm:"not null"`  //出价金额，单位为ETH或ERC20代币
	TxHash string `gorm:"type:varchar(66);not null"` //交易哈希
	BlockNumber uint64 `gorm:"not null"` //区块高度
}