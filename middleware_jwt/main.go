package main

import (
	"log"
	"middleware_jwt/handler"
	"middleware_jwt/middleware"
	"middleware_jwt/pkg"
	"middleware_jwt/sqlc"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	dbConn := pkg.NewDBConn(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))

	queries := sqlc.New(dbConn)

	authHandler := handler.NewAuthHandler(queries)
	userHandler := handler.NewUserHandler(queries)

	err := pkg.InitializeKeys()

	if err != nil {
		log.Fatal("err initialize keys: ", err.Error())
	}

	r := gin.New()
	r.Use(middleware.LoggingMiddleware())

	public := r.Group("/api")
	public.POST("/register", authHandler.Register)
	public.POST("/login", authHandler.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", userHandler.Profile)

	r.Run(":8080")
}
