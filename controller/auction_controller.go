package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"encoding/json"
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



//获取某个拍卖的出价列表
func(ac *AuctionController) GetBids(c *gin.Context){
	//获取拍卖id
	id := c.Param("id")

	var bids []model.Bid

	if err := ac.db.
				Where("auction_id = ?",id).
				Order("created_at asc").
				Find(&bids).Error;err != nil {
					c.JSON(500,gin.H{
						"error":err.Error()
					})
					return
				}

	c.JSON(200,bids)
}



//获取拍卖总数和出价总数
func(ac *AuctionController) GetPlatformStats(c *gin.Context){
	var auctionCount int64
	var bidCount int64

	ac.db.Table("auctions").Count(&auctionCount)
	ac.db.Table("bids").Count(&bidCount)

	c.JSON(http.StatusOK, gin.H{
		"totalAuctions":auctionCount,
		"totalBids":bidCount
	})
}


//根据钱包地址获取NFT
func GetNFTsByAddress(c *gin.Context){
	address := c.Param("address")

	url := fmt.Sprintf(
		"https://eth-sepolia.g.alchemy.com/nft/v3/%s/getNFTsForOwner?owner=%s",
		os.Getenv("ALCHEMY_API_KEY"),
		address,
	)

	resp,err := http.Get(url)

	if err != nil {
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	c.JSON(200,data)
}


func (ac *AuctionController) GetNFTsByAddressByMoralis(c *gin.Context) {
	address := c.Param("address")

	url := fmt.Sprintf(
		"https://deep-index.moralis.io/api/v2.2/%s/nft?chain=sepolia&format=decimal&limit=25",
		address,
	)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", os.Getenv("MORALIS_API_KEY"))
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(res.Body).Decode(&data)

	c.JSON(200, data)
}
