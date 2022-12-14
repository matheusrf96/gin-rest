package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusrf96/gin-rest/controllers"
)

func HandleRequests(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Aopa!",
		})
	})
	r.GET("/students", controllers.GetStudents)
	r.GET("/students/:id", controllers.GetStudentsById)
	r.GET("/students/cpf/:cpf", controllers.GetStudentsByCpf)
	r.POST("/students", controllers.CreateStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.RemoveStudent)
}
