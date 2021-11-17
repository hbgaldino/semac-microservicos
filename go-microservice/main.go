package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {

		hostname, _ := os.Hostname()
		c.JSON(200, gin.H{
			"hostname": hostname,
		})
	})
	r.Run()
}
