package main

import (
	"web-service-gin/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//homepage
	router.GET("/", service.HomePage)
	//information page about calculator
	router.GET("/calculator", service.Information)
	//post data user ke handler
	router.POST("/calculator", service.CalculatorHandler)

	//Server
	// fmt.Println("===Server Belajan di localhost:8080")
	router.Run("localhost:8080")
}
