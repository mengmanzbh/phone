package api

import (
	"github.com/gin-gonic/gin" 
    "fmt"
    "mobile/utils"
    "net/url"
    "encoding/json"
)
const APPKEY = "f17e30d841c5b17dbc00605a556d549a" //您申请的APPKEY
const OpenID = "JH24c5c602f7d064b5bd591d3eff1efe8e" //OpenID在个人中心查询
// 检测手机号码是否能充值
func Telcheck(ctx *gin.Context) {
    //获取请求参数
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
    data,err:=utils.Post(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
        ctx.JSON(404, gin.H{
            "code": "404",
            "message": err,
            })
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)

        ctx.JSON(200, gin.H{
            "error_code": netReturn["error_code"],
            "message": netReturn["reason"],
            "result":netReturn["result"],
        })
    }

}
// 根据手机号和面值查询商品
func Telquery(ctx *gin.Context) {
	phoneno := ctx.PostForm("phoneno")
    cardnum := ctx.PostForm("cardnum")
        //请求地址
    juheURL :="http://op.juhe.cn/ofpay/mobile/telquery"
 
    //初始化参数
    param:=url.Values{}
 
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("phoneno",phoneno) //手机号码
    param.Set("cardnum",cardnum) //充值金额,目前可选：5、10、20、30、50、100、300
    param.Set("key",APPKEY) //应用APPKEY(应用详细页查询)
 
 
    //发送请求
    data,err:=utils.Post(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
        ctx.JSON(404, gin.H{
            "code": "404",
            "message": err,
            })
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)

        ctx.JSON(200, gin.H{
            "error_code": netReturn["error_code"],
            "message": netReturn["reason"],
            "result":netReturn["result"],
        })
    }
}
// 手机直充接口
func Onlineorder(ctx *gin.Context) {
	phoneno := ctx.PostForm("phoneno")
    cardnum := ctx.PostForm("cardnum")
    orderid := utils.GetRandomString(6)
        //请求地址
    juheURL :="http://op.juhe.cn/ofpay/mobile/onlineorder"
 
    //初始化参数
    param:=url.Values{}
    
    //校验值，md5(OpenID+key+phoneno+cardnum+orderid)
    sign := utils.MD5(OpenID+APPKEY+phoneno+cardnum+orderid)

    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("phoneno",phoneno) //手机号码
    param.Set("cardnum",cardnum) //充值金额,目前可选：5、10、20、30、50、100、300
    param.Set("orderid",orderid) //商家订单号，8-32位字母数字组合
    param.Set("key",APPKEY) //应用APPKEY(应用详细页查询)
    param.Set("sign",sign) //校验值，md5(OpenID+key+phoneno+cardnum+orderid)，OpenID在个人中心查询
 
 
    //发送请求
    data,err:=utils.Post(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
        ctx.JSON(404, gin.H{
            "code": "404",
            "message": err,
            })
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)

        ctx.JSON(200, gin.H{
            "error_code": netReturn["error_code"],
            "message": netReturn["reason"],
            "result":netReturn["result"],
        })
    }

}
// 订单状态查询
func Ordersta(ctx *gin.Context) {
	// orderid := ctx.PostForm("orderid")
}
