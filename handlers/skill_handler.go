package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "Go-API/models"
    "Go-API/db"
)

func GetSkills(context *gin.Context) {
    var skills []models.Skill
    err := db.GetDB().Select(&skills, "SELECT * FROM skill")
    if err != nil {
        context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
        return
    }

    context.IndentedJSON(http.StatusOK, skills)
}

func GetSkill(context *gin.Context) {
    id := context.Param("id")
    var s models.Skill
    err := db.GetDB().Get(&s, "SELECT * FROM skill WHERE skillId=?", id)
    if err != nil {
        context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
        return
    }

    context.IndentedJSON(http.StatusOK, s)
}

func CreateSkill(context *gin.Context) {
    var newSkill models.Skill
    if err := context.BindJSON(&newSkill); err != nil {
        context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
        return
    }

    result, err := db.GetDB().Exec("INSERT INTO skill (category, skillName) VALUES (?, ?)", newSkill.Category, newSkill.SkillName)
    if err != nil {
        context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
        return
    }

    newSkill.SkillId, _ = result.LastInsertId()
    context.IndentedJSON(http.StatusCreated, newSkill)
}

func UpdateSkill(context *gin.Context) {
    id := context.Param("id")
    var currentSkill models.Skill
    err := db.GetDB().Get(&currentSkill, "SELECT * FROM skill WHERE skillId=?", id)
    if err != nil {
        context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
        return
    }

    var updatedSkill models.Skill
    if err := context.BindJSON(&updatedSkill); err != nil {
        context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
        return
    }

    _, err = db.GetDB().Exec("UPDATE skill SET category=?, skillName=? WHERE skillId=?", updatedSkill.Category, updatedSkill.SkillName, id)
    if err != nil {
        context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
        return
    }

    currentSkill.Category = updatedSkill.Category
    currentSkill.SkillName = updatedSkill.SkillName

    context.IndentedJSON(http.StatusOK, currentSkill)
}

func DeleteSkill(context *gin.Context) {
    id := context.Param("id")
    result, err := db.GetDB().Exec("DELETE FROM skill WHERE skillId=?", id)
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