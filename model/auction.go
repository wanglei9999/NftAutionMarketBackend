package model

import(
	"time"
	"gorm.io/gorm"
)

type Auction struct{
	gorm.Model
	AuctionID uint64 `gorm:"uniqueIndex;not null"` //拍卖ID，唯一标识一个拍卖
	SellerAddress string `gorm:"type:char(42);index;not null"` //卖家地址
	Duration uint64 `gorm:"not null"` //拍卖持续时间，单位为秒
	StartTime time.Time `gorm:"index;not null"` //拍卖开始时间
	EndTime time.Time `gorm:"index;not null"` //拍卖结束时间，等于StartTime + Duration
	StartPrice string `gorm:"type:varchar(64);not null"`  //起拍价
	HighestBid string `gorm:"type:varchar(64)"` //当前最高出价
	HighestBidder string `gorm:"type:char(42)"` //当前最高出价者
	NftContractAddress string `gorm:"type:char(42);index;not null"` //NFT合约地址
	TokenID string `gorm:"type:varchar(64);index;not null"` //NFT Token ID
	Ended bool `gorm:"index;not null"` //是否结束
	TokenAddress string `gorm:"type:char(42);index;not null"` //支付代币地址，0地址表示使用ETH
	
	}
