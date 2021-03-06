package infra

import (
	"assignment/infra/api"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router = gin.Default()

func InitRouter() {
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	Router.GET("/view/get_products", api.GetProducts)
	Router.GET("/view/get_users", api.GetUsers)

	Router.PATCH("/buy/:product/:user/:point", api.BuyProduct)
	Router.PATCH("/activity/:state", api.ChangeActivity)
}
