package main

import (
	"github.com/Asylniet/golang_lab3/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDb()

	r := gin.Default()
	setUpRoutes(r)
	r.Run(":3000")
}
