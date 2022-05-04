package controllers

import (
	"net/http"

	"github.com/Johnman67112/gin-rest-api/database"
	"github.com/Johnman67112/gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

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
