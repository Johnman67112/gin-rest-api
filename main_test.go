package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Johnman67112/gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func RoutesSetup() *gin.Engine {
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
