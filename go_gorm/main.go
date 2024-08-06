package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	_ "go_gorm/docs"
)

// @title Your API Title
// @version 1.0
// @description Your API Description
// @host localhost:8080
// @BasePath /api/v1
func main() {
	LoadEnv()
	InitDB()

	r := gin.Default()

	// Middleware
	//r.Use(AuthMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		users.Use()
		{
			users.GET("/", GetUsers)
			users.POST("/", CreateUser)
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}

	// Clean Db
	err := DB.Migrator().DropTable(&User{})
	if err != nil {
		return
	}
}
