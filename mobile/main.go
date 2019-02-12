package main

import (
	"mobile/api"

	_ "phone/docs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("/mobile/telcheck", api.Telcheck)
	r.POST("/mobile/telquery", api.Telquery)
	r.POST("/mobile/onlineorder", api.Onlineorder)
	r.POST("/mobile/ordersta", api.Ordersta)

	r.Run()
}
