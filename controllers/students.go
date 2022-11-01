package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusrf96/gin-rest/db"
	"github.com/matheusrf96/gin-rest/models"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	db.DB.Find(&students)

	c.JSON(http.StatusOK, students)
}

func GetStudentsById(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")
	db.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
