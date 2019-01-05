package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "io/ioutil"
    "net/url"
    "fmt"
    "encoding/json"
)
const APPKEY = "f17e30d841c5b17dbc00605a556d549a" //您申请的APPKEY
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
    data,err:=Post(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
        ctx.JSON(404, gin.H{
            "code": "404",
            "message": err,
            })
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)
        if netReturn["error_code"].(float64)==0{
            fmt.Printf("接口返回reason字段是:\r\n%v",netReturn["reason"])
            //返回给前端
            ctx.JSON(200, gin.H{
            "code": "200",
            "message": netReturn["reason"],
            })

        }else{
           ctx.JSON(400, gin.H{
            "code": "400",
            "message": netReturn["reason"],
            })

        }
    }

}
// 根据手机号和面值查询商品
func Telquery(ctx *gin.Context) {
	phoneno := ctx.PostForm("phoneno")
    cardnum := ctx.PostForm("cardnum")
}
// 手机直充接口
func Onlineorder(ctx *gin.Context) {
	phoneno := ctx.PostForm("phoneno")
    cardnum := ctx.PostForm("cardnum")
    orderid := ctx.PostForm("orderid")
    sign := ctx.PostForm("sign")
}
// 订单状态查询
func Ordersta(ctx *gin.Context) {
	orderid := ctx.PostForm("orderid")
}
// get 网络请求
func Get(apiURL string,params url.Values)(rs[]byte ,err error){
    var Url *url.URL
    Url,err=url.Parse(apiURL)
    if err!=nil{
        fmt.Printf("解析url错误:\r\n%v",err)
        return nil,err
    }
    //如果参数中有中文参数,这个方法会进行URLEncode
    Url.RawQuery=params.Encode()
    resp,err:=http.Get(Url.String())
    if err!=nil{
        fmt.Println("err:",err)
        return nil,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
 
// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values)(rs[]byte,err error){
    resp,err:=http.PostForm(apiURL, params)
    if err!=nil{
        return nil ,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
