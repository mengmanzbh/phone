package api
import (
	"github.com/gin-gonic/gin" 
    "fmt"
    "mobile/utils"
    "net/url"
    "encoding/json"
)

// 全部流量套餐列表
func Flowlist(ctx *gin.Context) {
    //请求地址
    juheURL :="http://v.juhe.cn/flow/list"
    //初始化参数
    param:=url.Values{}
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
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