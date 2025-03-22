package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/task", func(c *gin.Context) {
		taskId := c.Query("id")
	})
}
