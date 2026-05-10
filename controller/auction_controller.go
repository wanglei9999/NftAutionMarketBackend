package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gorm.io/gorm"
	"NftAutionMarketBackend/model"
	"NftAutionMarketBackend/database"
)


func NewAuctionController() *AuctionController {
	return &AuctionController{
		db: database.DB,
	}
}

type AuctionController struct {
	db *gorm.DB
}


//获取拍卖列表
func (ac *AuctionController) GetAuctions(c *gin.Context) {
	var auctions []model.Auction
	order := c.DefaultQuery("order","start_time desc")
	ended := c.DefaultQuery("ended","false")
	query := ac.db.Table("auctions")
	if ended == "true" {
		query = query.Where("ended = ?", true)
	}else {
		query = query.Where("ended = ?", false)
	}
	if err := query.Order(order).Find(&auctions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch auctions"})
		return
	}
	c.JSON(http.StatusOK, auctions)
	
}