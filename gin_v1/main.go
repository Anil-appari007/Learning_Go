package main

import (
	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.String(200, "HelloWorld")
}
func main() {
	router := gin.Default()
	router.GET("/", helloWorld)

	router.Run(":8000")
}
