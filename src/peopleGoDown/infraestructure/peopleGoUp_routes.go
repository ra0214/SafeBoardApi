package infraestructure

import (
	"apiMulti/src/peopleGoDown/application"
	"apiMulti/src/peopleGoDown/domain"
	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IPeopleGoDown, rabbitRepo domain.IPeopleGoDownRabbitqm) *gin.Engine {
	r := gin.Default()

	CreatePeopleGoDown := application.NewCreatepeopleGoDown(rabbitRepo, repo)
	createPeopleGoDownController := NewCreatePeopleGoDownController(CreatePeopleGoDown)

	viewPeopleGoDown := application.NewViewPeopleGoDown(repo)
	viewPeopleGoDownController := NewViewPeopleGoDownController(viewPeopleGoDown)

	r.POST("/peopleGoDown", createPeopleGoDownController.Execute)
	r.GET("/peopleGoDown", viewPeopleGoDownController.Execute)

	return r
}