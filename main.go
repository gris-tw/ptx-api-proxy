package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/proxy-api", func(c *gin.Context) {
		c.String(200, "Proxy is working.")
	})

	r.POST("/ptx-service", func(c *gin.Context) {

		apiurl := c.PostForm("apiString")
		req, err := http.NewRequest("GET", apiurl, nil)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		c.String(200, string(body))
	})

	r.Run(":80")
}
