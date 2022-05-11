package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestSearchStudentByID(t *testing.T) {
	database.DatabaseConnect()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := RoutesSetup()
	r.GET("/students/:id", controllers.GetStudent)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var studentMock models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentMock)

	assert.Equal(t, "Student's Name Test", studentMock.Name, "Names should be the same")
	assert.Equal(t, "123456789", studentMock.RG, "RGs should be the same")
	assert.Equal(t, "12345678901", studentMock.CPF, "CPFs should be the same")
	assert.Equal(t, http.StatusOK, resp.Code, "Status code should be OK")
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

func TestDeleteStudent(t *testing.T) {
	database.DatabaseConnect()
	CreateStudentMock()

	r := RoutesSetup()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	path := "/students/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", path, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestEditStudent(t *testing.T) {
	database.DatabaseConnect()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := RoutesSetup()
	r.PATCH("/students/:id", controllers.EditStudent)

	student := models.Student{Name: "Student's Name Test", CPF: "47123456789", RG: "123456700"}
	valueJson, _ := json.Marshal(student)
	path := "/students/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(valueJson))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var studentMock models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentMock)

	assert.Equal(t, "Student's Name Test", studentMock.Name, "Names should be the same")
	assert.Equal(t, "123456700", studentMock.RG, "RGs should have been updated")
	assert.Equal(t, "47123456789", studentMock.CPF, "CPFs should have been updated")
	assert.Equal(t, http.StatusOK, resp.Code, "Status code should be OK")
}
