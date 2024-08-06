package main

import (
	"consume_api/rest"
	"consume_api/soap"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var restClient *rest.RClient
var soapClient *soap.SClient

func main() {

	restUrl := "https://jsonplaceholder.typicode.com/posts"
	soapUrl := "http://www.dneonline.com/calculator.asmx"

	restClient = rest.NewRESTClient(restUrl, 5*time.Second)
	soapClient = soap.NewSOAPClient(soapUrl, 5*time.Second)

	server := gin.Default()

	server.GET("/rest", getAllPosts)

	server.POST("/soap", calculateMultiplication)

	err := server.Run(":8080")
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
