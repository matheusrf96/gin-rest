package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrf96/gin-rest/controllers"
)

func HandleRequests(r *gin.Engine) {
	r.GET("/students", controllers.GetStudents)
	r.GET(":name", controllers.Aopa)
	r.POST("/students", controllers.CreateStudent)
}
