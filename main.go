package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrf96/gin-rest/db"
	"github.com/matheusrf96/gin-rest/routes"
)

func main() {
	db.ConnectDatabase()

	r := gin.Default()

	routes.HandleRequests(r)

	r.Run(":5000")
}
