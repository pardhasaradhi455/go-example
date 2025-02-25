package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"github.com/gin-gonic/gin"
)

var code int

func init(){
	code = rand.Intn(100)
}

func router() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		fmt.Println("the default end point '/'")
		ctx.JSON(http.StatusOK, "the default end point")
	})
	router.GET("/message", func(ctx *gin.Context) {
		fmt.Println("triggered end point '/health'")
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "your application is deployed successfully",
			"code" : code,
		})
	})

	router.Run()
}