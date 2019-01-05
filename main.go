package main

import (
	"phone/api"

	_ "phone/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	r := gin.New()
	r.POST("/mobile/telcheck", api.Telcheck)
	r.POST("/mobile/telquery", api.Telquery)
	r.POST("/mobile/onlineorder", api.Onlineorder)
	r.POST("/mobile/ordersta", api.Ordersta)

	r.Run()
}
