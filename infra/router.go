package infra

import (
	"server/infra/api"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router = gin.Default()

func InitRouter() {
	api.Init()

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.GET("/buy/:product/:vip_level", api.GetGroups)
	Router.GET("/group/get_by_name/:group_name", api.GetGroupByName)

	Router.POST("/group/create/:group_name", api.CreateGroup)

	Router.PATCH("/group/update/:search_name/:set_name", api.UpdateGroupName)

	Router.GET("/books", api.GetBooks)
	Router.GET("/books/:id", api.BookById)

	Router.POST("/books", api.CreateBook)

	Router.PATCH("/checkout", api.CheckoutBook)
	Router.PATCH("/return", api.ReturnBook)
}
