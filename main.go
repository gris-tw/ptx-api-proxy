package main

import (
	"github.com/gin-gonic/gin"

	"github.com/hpcslag/Sample-code/Golang/lib"
)

func main() {

	ptxapi := lib.PTXService{
		AppID:  "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF",
		AppKey: "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF",
	}

	r := gin.Default()
	r.GET("/proxy-api", func(c *gin.Context) {
		c.String(200, "Proxy is working.")
	})

	r.POST("/ptx-service", func(c *gin.Context) {

		apiurl := c.PostForm("apiString")

		response := ptxapi.Get(apiurl)

		c.String(200, response)
	})

	r.Run(":80")
}
