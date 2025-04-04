package infraestructure

import (
	"apiMulti/src/peopleGoDown/application"
	"apiMulti/src/peopleGoDown/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IPeopleGoDown, rabbitRepo domain.IPeopleGoDownRabbitqm) *gin.Engine {
	r := gin.Default()

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
