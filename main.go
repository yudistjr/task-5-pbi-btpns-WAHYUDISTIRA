package main

import (
	"user-personalize/database"
	"user-personalize/router"

	"github.com/gin-gonic/gin"
)

func init() {
	database.Connect()
	database.Migrate()
}

func main() {
	r := gin.Default()
	router.Routing(r)
	r.Run()
}