package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type skill struct {
	SkillId   string `json:"skillId"`
	Category  string `json:"category"`
	SkillName string `json:"skillName"`
}

var skills = []skill{
	{SkillId: "1", Category: "Programming Languages", SkillName: "C#"},
	{SkillId: "2", Category: "Programming Languages", SkillName: "Java"},
	{SkillId: "3", Category: "Programming Languages", SkillName: "Python"},
}

func getSkills(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, skills)
}

func getSkill(context *gin.Context) {
	id := context.Param("id")
	skill, err := getSkillBySkillId(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, skill)
}

func createSkill(context *gin.Context) {
	var newSkill skill
	if err := context.BindJSON(&newSkill); err != nil {
		return
	}
	skills = append(skills, newSkill)
	context.IndentedJSON(http.StatusCreated, newSkill)
}

func updateSkill(context *gin.Context) {
	id := context.Param("id")
	currentskill, err := getSkillBySkillId(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	var updatedSkill skill
	if err := context.BindJSON(&updatedSkill); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	currentskill.Category = updatedSkill.Category
	currentskill.SkillName = updatedSkill.SkillName

	context.IndentedJSON(http.StatusOK, currentskill)
}

func deleteSkill(context *gin.Context) {
	id := context.Param("id")
	index, err := findSkillIndex(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	skills = append(skills[:index], skills[index+1:]...)

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Skill deleted successfully"})
}

func getSkillBySkillId(skillId string) (*skill, error) {
	for i, s := range skills {
		if s.SkillId == skillId {
			return &skills[i], nil
		}
	}

	return nil, errors.New("skill not found")
}

func findSkillIndex(skillId string) (int, error) {
	for i, s := range skills {
		if s.SkillId == skillId {
			return i, nil
		}
	}

	return -1, errors.New("skill not found")
}

func main() {
	router := gin.Default()
	router.GET("/skills", getSkills)
	router.GET("/skill/:id", getSkill)
	router.POST("/skill", createSkill)
	router.PUT("/skill/:id", updateSkill)
	router.DELETE("/skill/:id", deleteSkill)
	router.Run("localhost:8080")
}
