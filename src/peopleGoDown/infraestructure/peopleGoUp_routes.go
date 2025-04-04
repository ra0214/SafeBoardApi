package infraestructure

import (
	"apiMulti/src/peopleGoDown/application"
	"apiMulti/src/peopleGoDown/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IPeopleGoDown, rabbitRepo domain.IPeopleGoDownRabbitqm) *gin.Engine {
	r := gin.Default()

	// Configurar CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Casos de uso
	createPeopleGoDown := application.NewCreatePeopleGoDown(rabbitRepo, repo)
	viewPeopleGoDown := application.NewViewPeopleGoDown(repo)

	// Controladores
	createPeopleGoDownController := NewCreatePeopleGoDownController(createPeopleGoDown)
	viewPeopleGoDownController := NewViewPeopleGoDownController(viewPeopleGoDown)

	// Rutas
	r.POST("/peopleGoDown", createPeopleGoDownController.Execute)
	r.GET("/peopleGoDown", viewPeopleGoDownController.Execute)

	return r
}
