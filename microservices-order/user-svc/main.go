package main

import (
	"log"
	"user-svc/internal/app"
	"user-svc/pkg"

	"github.com/gin-gonic/gin"
)

func main() {

	config := pkg.LoadConfig()

	log.Println("config: ", config)

	gin := gin.New()
	app.NewApp(gin, config).Run()
}
