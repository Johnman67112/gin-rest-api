package routes

import (
	"github.com/Johnman67112/gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.GetStudents)
	r.GET("/:name", controllers.Hello)
	r.GET("/students/:id", controllers.GetStudent)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)

	r.POST("/students", controllers.CreateStudent)

	r.DELETE("/students/:id", controllers.DeleteStudent)

	r.PATCH("/students/:id", controllers.EditStudent)

	r.Run()
}
