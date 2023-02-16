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
	router.DELETE("/procadd/:id", controller.DeleteAddressController)

	router.GET("/propvals/", controller.GetPropertyValsController)
	router.GET("/propvals/:id", controller.GetIDPropertyvalsController)
	router.POST("/propvals/", controller.PostPropertyvalsController)
	router.PUT("/propvals/:id", controller.EditPropertyvalsController)
	router.DELETE("/propvals/:id", controller.DeletePropertyvalsController)

	router.POST("/filter/", controller.FilterBased)

}
