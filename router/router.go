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

	}

}



