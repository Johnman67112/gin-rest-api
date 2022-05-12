package routes

import (
	"github.com/Johnman67112/gin-rest-api/controllers"
	"github.com/gin-gonic/gin"

	docs "github.com/Johnman67112/gin-rest-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	docs.SwaggerInfo.BasePath = "/"

	r.GET("/:name", controllers.Hello)

	r.GET("/students", controllers.GetStudents)

	r.GET("/students/:id", controllers.GetStudent)

	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)

	r.GET("/index", controllers.ShowIndex)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.NoRoute(controllers.RouteNotFound)

	r.POST("/students", controllers.CreateStudent)

	r.DELETE("/students/:id", controllers.DeleteStudent)

	r.PATCH("/students/:id", controllers.EditStudent)

	r.Run()
}
