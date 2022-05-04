package main

import (
	"github.com/Johnman67112/gin-rest-api/database"
	"github.com/Johnman67112/gin-rest-api/routes"
)

func main() {
	database.DatabaseConnect()
	routes.HandleRequests()
}
