package controllers

import (
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
