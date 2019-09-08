package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	fmt.Println("Hello World")
	serverSetup()
	r.Run(":3000")
}

func serverSetup() {
	r = gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Origin", "Content-Length", "Content-Type"},
		AllowOrigins:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	gin.SetMode(gin.ReleaseMode)
	//Serve the files here:

	r.GET("/hello", HelloWorld)
	r.GET("/", MainWorld)
	// setup route group
}

//HelloWorld is an exported function. Comment is right here
func HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "you get ahello newewew")
}

//MainWorld is an exported function. Comment is right here
func MainWorld(c *gin.Context) {
	c.String(http.StatusOK, "main payge yea....")
}
