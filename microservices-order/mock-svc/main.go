package main

import (
	"fmt"
	"mock-svc/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	gin := gin.Default()

	uh := handler.NewUserHandler()
	ph := handler.NewProductHandler()
	pyh := handler.NewPaymentHandler()

	gin.GET("/users", uh.GetUser)
	gin.GET("/users/:id", uh.GetUserByID)

	gin.GET("/products", ph.GetProduct)
	gin.POST("/products/reserve", ph.ReserveProduct)
	gin.POST("/products/release", ph.ReleaseProduct)

	gin.GET("/balance", pyh.GetBalance)
	gin.POST("/payment", pyh.CreateTransaction)
	gin.POST("/refund", pyh.RefundTransaction)

	fmt.Println("Server is running on port 5000")
	gin.Run(":5000")
}
