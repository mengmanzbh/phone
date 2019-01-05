package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
	// "io/ioutil"
)

// @Summary 检测手机号码是否能充值
// @Description 检查手机号码是否能充值
// @ID phoneno
// @Accept json
// @Produce json
// @Param phoneno path int true "phoneno"
// @Param cardnum body web.Pet true "cardnum"
// @Success 200 {string} string "手机可以充值"
// @Failure 400 {string} web.APIError "手机号码必须输入"
// @Router /phoneapi/Telcheck [POST]
func Telcheck(ctx *gin.Context) {
	 // phoneno	是	string	手机号码
     // cardnum	是	string	充值金额,目前可选：1、2、5、10、20、30、50、100、200、300、500
     // key	是	string	在个人中心->我的数据,接口名称上方查看

	    phoneno := ctx.PostForm("phoneno")
	    cardnum := ctx.PostForm("cardnum")
	    key := ctx.PostForm("key")

        ctx.JSON(200, gin.H{
            "status":  "posted",
            "phoneno": phoneno,
            "cardnum":    cardnum,
            "key":    key,
        })

}

// @Description get struct array by ID
// @ID get-struct-array-by-string
// @Accept json
// @Produce json
// @Param some_id path string true "Some ID"
// @Param offset query int true "Offset"
// @Param limit query int true "Offset"
// @Success 200 {string} string "ok"
// @Failure 400 {object} web.APIError "We need ID!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetStructArrayByString.")
}

// @Summary Upload file
// @Description Upload file
// @ID file.upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "this is a test file"
// @Success 200 {string} string "ok"
// @Failure 400 {object} web.APIError "We need ID!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /file/upload [post]
func Upload(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Upload")
}

// @Summary use Anonymous field
// @Success 200 {object} web.RevValue "ok"
func AnonymousField() {

}

type Pet3 struct {
	ID int `json:"id"`
}
