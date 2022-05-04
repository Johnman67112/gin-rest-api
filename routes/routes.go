package routes

import (
	"github.com/Johnman67112/gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetStudents)
	r.GET("/:name", controllers.Hello)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.GetStudent)
	r.Run()
}
