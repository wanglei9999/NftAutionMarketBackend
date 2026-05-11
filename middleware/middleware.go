package middleware

import (
	"net/http"
	"rentime/debug"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"NftAutionMarketBackend/auth"
)


//校验token
func AutoMiddleware() gin.HandlerFunc{

	return func(ctx *gin.Context){
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == ""{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error" : "Missing token"
			})
		}
		
		tokenStr := strings.TrimPrefix(authHeader,"Bearer ")
		token,err := jwt.Parse(tokenStr,func(t *jwt.Token)(any,error){
			return auth.JwtSecret,nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error":"Invalid token"})
			return
		}
		//转换成map
		claims,ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error":"Invalid token"})
			return
		}

		address := claims["address"]


		//将地址放入上下文
		ctx.Set("address",address)

		//放行
		ctx.Next()
	}

}


//全局日志打印
func RecoveryHandler() gin.HandlerFunc {
	return func(c *gin.Context){
		defer func(){
			if err := recover(); err != nil{
				//打印错误堆栈
				debug.PrintStack()

				//自定义返回
				c.JSON(http.StatusInternalServerError,gin.H{
					"code": 1,
					"message":"服务器内部错误，请稍后再试",
				})
				//阻止继续执行
				c.Abort()
			}
		}()
		//继续处理请求
		c.Next()
	}
}
