package main

import (
	"net/http"
	"os"
	"Go-API/db"
	"Go-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()

	routes.SkillRoutes(router);

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		panic(err)
	}
}