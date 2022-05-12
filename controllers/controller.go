package controllers

import (
	"net/http"

	"github.com/Johnman67112/gin-rest-api/database"
	"github.com/Johnman67112/gin-rest-api/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

//Hello godoc
// @Summary      Says hello
// @Description  Says hello to user
// @Tags         hello
// @Accept       json
// @Produce      json
// @Param        name  path  string  true  "Name"
// @Success      200  {object}  string
// @Failure      400  {object}  httputil.HTTPError
// @Router       /{name} [get]
func Hello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API said": "What's up " + name + ", how u doing?",
	})
}

//CreateStudent godoc
// @Summary      Creates a new student
// @Description  With params Name, RG, CPF creates a new student
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        student  body  models.Student  true  "Student Model"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Router       /students [post]
func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.StudentValidator(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}

//GetStudents godoc
// @Summary      Show all students
// @Description  Route to show all students
// @Tags         students
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Router       /students [get]
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

func GetStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

//Delete
func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{"data": "Student deleted sucessfully"})
}

//Edit
func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.StudentValidator(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func ShowIndex(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
