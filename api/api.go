package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
	// "io/ioutil"
)

// @Summary 检测手机号码是否能充值
// @Description get string by ID
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} web.APIError "We need ID!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
func Telcheck(ctx *gin.Context) {

	    username := ctx.PostForm("username")
	    password := ctx.PostForm("password")

        ctx.JSON(200, gin.H{
            "status":  "posted",
            "message": username,
            "nick":    password,
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
