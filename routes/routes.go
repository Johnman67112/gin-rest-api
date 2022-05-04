package routes

import (
	"github.com/Johnman67112/gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowStudents)
	r.GET("/:name", controllers.Hello)
	r.POST("/students", controllers.CreateStudent)
	r.Run()
}
