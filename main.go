package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.GET("/proxy-api", func(c *gin.Context){
		c.String(200, "Proxy is working.")
	})
}