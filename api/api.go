package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
	// "io/ioutil"
)

// 检测手机号码是否能充值
func Telcheck(ctx *gin.Context) {

	    phoneno := ctx.PostForm("phoneno")
	    cardnum := ctx.PostForm("cardnum")
	    key := ctx.PostForm("key")

        ctx.JSON(200, gin.H{
            "status": "posted",
            "phoneno": phoneno,
            "cardnum": cardnum,
            "key":key,
        })

}
// 根据手机号和面值查询商品
func Telquery(ctx *gin.Context) {
	
}
// 手机直充接口
func Onlineorder(ctx *gin.Context) {
	
}
// 订单状态查询
func Ordersta(ctx *gin.Context) {
	
}
