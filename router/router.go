package router


import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"NftAutionMarketBackend/controller"
	"NftAutionMarketBackend/middleware"
)



func InitRouter() *gin.Engine{
	r := gin.New()
	r.Use(
		//日志打印格式中间件
		gin.LoggerWithFormatter(
			func(param gin.LogFormatterParams) string  {
				return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
				)
			}
		),
		middleware.RecoveryHandler(),
	)


	//用户控制器
	userController := controller.NewUserController()



	authGroup := r.Group("/auth")
	{
		//获取nonce请求
		authGroup.Post("/nonce",userController.GetNonceHandler)

		//登录获取token
		auth.Post("login",userController.LoginHandler)

	}



	//拍卖控制器
	auctionController := controller.NewAuctionController()
    apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware){
		//获取拍卖列表
		api.GET("/auctions", auctionController.GetAuctions)
		//获取某个拍卖的出价列表
		api.GET("/auctions/:id/bids", auctionController.GetAuctionBids)
		//获取拍卖总数和出价总数
		api.GET("/stats", auctionController.GetPlatformStats)
		//获取某个用户地址拥有的NFT
		api.GET("/nfts/:address", auctionController.GetNFTsByAddress)
		api.GET("/nfts/moralis/:address", auctionController.GetNFTsByAddressByMoralis)


	}

	return r

}



