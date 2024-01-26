package main

import (
	"net/http"
	"os"
	"Go-API/db"
	"Go-API/routes"
)

func main() {
	db.InitDB()

	router := routes.SkillRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		panic(err)
	}
}