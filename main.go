package main

import (
	"github.com/gin-gonic/gin"
	"github.com/npinnaka/goproject/db"
	"github.com/npinnaka/goproject/routes"
)

func main() {
	server := gin.Default()
	db.InitDB()
	routes.RegisterRoutes(server)
	server.Run(":8080")
	defer db.CloseDB()
}
