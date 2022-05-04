package controllers

import (
	"net/http"

	"github.com/Johnman67112/gin-rest-api/database"
	"github.com/Johnman67112/gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API said:": "What's up " + name + ", how u doing?",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func GetStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}
