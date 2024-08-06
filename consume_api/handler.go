package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type MultiplyRequest struct {
	XMLName xml.Name `xml:"Multiply"`
	Xmlns   string   `xml:"xmlns,attr"`
	IntA    int      `xml:"intA"`
	IntB    int      `xml:"intB"`
}

type MultiplyResponse struct {
	XMLName        xml.Name `xml:"MultiplyResponse"`
	MultiplyResult int      `xml:"MultiplyResult"`
}

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

func getAllPosts(c *gin.Context) {

	var posts []Post

	err := restClient.Call(nil, &posts)
	if err != nil {
		log.Println("REST API error: ", err.Error())
		c.JSON(500, gin.H{"error": err.Error})
		return
	}
	c.JSON(200, posts)
}

func calculateMultiplication(c *gin.Context) {

	var request Request
	var soapResponse MultiplyResponse

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("SOAP API error: ", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	mRequest := MultiplyRequest{
		Xmlns: "http://tempuri.org/",
		IntA:  request.A,
		IntB:  request.B,
	}

	err = soapClient.Call("http://tempuri.org/Multiply", mRequest, &soapResponse)
	if err != nil {
		log.Println("SOAP API error: ", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, soapResponse)
}
