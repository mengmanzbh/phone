package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
	// "io/ioutil"
)
const APPKEY = "f17e30d841c5b17dbc00605a556d549a" //您申请的APPKEY
// 检测手机号码是否能充值
func Telcheck(ctx *gin.Context) {

	phoneno := ctx.PostForm("phoneno")
    //请求地址
    juheURL :="http://op.juhe.cn/ofpay/mobile/telcheck"
    //初始化参数
    param:=url.Values{}
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("phoneno",phoneno) //手机号码
    param.Set("cardnum","5") //充值金额,目前可选：5、10、20、30、50、100、300
    param.Set("key",APPKEY) //应用APPKEY(应用详细页查询)
 
 
    //发送请求
    data,err:=Get(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
        ctx.JSON(400, gin.H{
            "code": "400",
            "message": err,
            })
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)
        if netReturn["error_code"].(float64)==0{
            fmt.Printf("接口返回reason字段是:\r\n%v",netReturn["reason"])

            ctx.JSON(200, gin.H{
            "code": "200",
            "message": netReturn["reason"],
            })

        }else{
           ctx.JSON(200, gin.H{
            "code": "200",
            "message": netReturn["reason"],
            })
        	
        }
    }

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
