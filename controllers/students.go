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

func GetStudentsByCpf(c *gin.Context) {
	var student models.Student

	cpf := c.Param("cpf")
	db.DB.Where(&models.Student{CPF: cpf}).First(&student)

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

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	db.DB.First(&student, id)

	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func RemoveStudent(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")
	db.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student was deleted succesfully",
	})
}
