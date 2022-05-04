package main

import (
	"github.com/Johnman67112/gin-rest-api/database"
	"github.com/Johnman67112/gin-rest-api/models"
	"github.com/Johnman67112/gin-rest-api/routes"
)

func main() {
	database.DatabaseConnect()
	models.Students = []models.Student{
		{Name: "Johnny Rockets", CPF: "00000000000", RG: "000000000"},
		{Name: "McDonalds", CPF: "11111111111", RG: "111111111"},
	}
	routes.HandleRequests()
}
