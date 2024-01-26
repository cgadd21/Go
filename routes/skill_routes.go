package routes

import (
	"Go-API/handlers"

	"github.com/gin-gonic/gin"
)

func SkillRoutes() *gin.Engine {
	router := gin.Default()

	skillGroup := router.Group("/skill")
	{
		skillGroup.GET("", handlers.GetSkills)
		skillGroup.GET("/:id", handlers.GetSkill)
		skillGroup.POST("", handlers.CreateSkill)
		skillGroup.PUT("/:id", handlers.UpdateSkill)
		skillGroup.DELETE("/:id", handlers.DeleteSkill)
	}

	return router
}
