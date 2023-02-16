package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karthiksbh/go-lang-task/controller"
)

func AddressRoute(router *gin.Engine) {
	router.GET("/procadd/", controller.GetAddressController)
	router.GET("/procadd/:id", controller.GetIDAddressController)
	router.POST("/procadd/", controller.PostAddressController)
	router.PUT("/procadd/:id", controller.EditAddressController)
	router.POST("/procadd/:id", controller.DeleteAddressController)
}
