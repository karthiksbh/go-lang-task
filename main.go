package main

import (
	"github.com/gin-gonic/gin"
	"github.com/karthiksbh/go-lang-task/config"
	"github.com/karthiksbh/go-lang-task/routes"
)

func main() {
	router := gin.New()
	config.Connectdb()
	routes.AddressRoute(router)
	router.Run(":8080")
}
