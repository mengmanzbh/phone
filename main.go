package main

import (
	"phone/controller"
	_ "phone/docs"
	"phone/httputil"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// @title Swagger 示例API
// @version 1.0
// @description 这是一个http server示例，用来演示swagger文档
// @termsOfService API服务条款:http://swagger.io/terms

// @contact.name 公开的API的联系信息:
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @licenes.name Apache 2.0
// @licenes.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

//@securityDefinitions.basic.BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydeginitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth.token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
			accounts.POST("", c.AddAccount)
			accounts.DELETE(":id", c.DeleteAccount)
			accounts.PATCH(":id", c.UpdateAccount)
			accounts.POST(":id/images", c.UploadAccountImage)
		}

	}
	// http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

func auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(context.GetHeader("Authorization")) == 0 {
			httputil.NewError(context, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			context.Abort()
		}
		context.Next()
	}
}
