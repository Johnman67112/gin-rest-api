package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Johnman67112/gin-rest-api/controllers"
	"github.com/Johnman67112/gin-rest-api/database"
	"github.com/Johnman67112/gin-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func RoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func TestVerifyStatusCodeFromHello(t *testing.T) {
	r := RoutesSetup()
	r.GET("/:name", controllers.Hello)

	req, _ := http.NewRequest("GET", "/johnny", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Should be the same")

	respMock := `{"API said":"What's up johnny, how u doing?"}`
	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, respMock, string(respBody), "Should be the same")
}

func CreateStudentMock() {
	student := models.Student{Name: "Student's Name Test", CPF: "12345678901",
		RG: "123456789"}

	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestListAllStudents(t *testing.T) {
	database.DatabaseConnect()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := RoutesSetup()
	r.GET("/students", controllers.GetStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestSeachStudentByCPF(t *testing.T) {
	database.DatabaseConnect()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := RoutesSetup()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)

	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
