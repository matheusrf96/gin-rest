package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   "1",
		"name": "matheus",
	})
}

func main() {
	r := gin.Default()

	r.GET("/students", getStudents)

	r.Run(":5000")
}
