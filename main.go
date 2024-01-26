package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type skill struct {
	SkillId   int64  `json:"skillId" db:"skillId"`
	Category  string `json:"category" db:"category"`
	SkillName string `json:"skillName" db:"skillName"`
}

var db *sqlx.DB

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := user + ":" + password + "@tcp(" + host + ")/" + dbName
   	db, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func getSkills(context *gin.Context) {
	var skills []skill
	err := db.Select(&skills, "SELECT * FROM skill")
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, skills)
}

func getSkill(context *gin.Context) {
	id := context.Param("id")
	var s skill
	err := db.Get(&s, "SELECT * FROM skill WHERE skillId=?", id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, s)
}

func createSkill(context *gin.Context) {
	var newSkill skill
	if err := context.BindJSON(&newSkill); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	result, err := db.Exec("INSERT INTO skill (category, skillName) VALUES (?, ?)", newSkill.Category, newSkill.SkillName)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	newSkill.SkillId, _ = result.LastInsertId()
	context.IndentedJSON(http.StatusCreated, newSkill)
}

func updateSkill(context *gin.Context) {
	id := context.Param("id")
	var currentSkill skill
	err := db.Get(&currentSkill, "SELECT * FROM skill WHERE skillId=?", id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	var updatedSkill skill
	if err := context.BindJSON(&updatedSkill); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	_, err = db.Exec("UPDATE skill SET category=?, skillName=? WHERE skillId=?", updatedSkill.Category, updatedSkill.SkillName, id)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	currentSkill.Category = updatedSkill.Category
	currentSkill.SkillName = updatedSkill.SkillName

	context.IndentedJSON(http.StatusOK, currentSkill)
}

func deleteSkill(context *gin.Context) {
	id := context.Param("id")
	result, err := db.Exec("DELETE FROM skill WHERE skillId=?", id)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Skill deleted successfully"})
}

func main() {
	initDB()

	router := gin.Default()
	router.GET("/skills", getSkills)
	router.GET("/skill/:id", getSkill)
	router.POST("/skill", createSkill)
	router.PUT("/skill/:id", updateSkill)
	router.DELETE("/skill/:id", deleteSkill)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
