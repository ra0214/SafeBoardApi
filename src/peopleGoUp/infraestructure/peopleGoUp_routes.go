package infraestructure

import (
	"apiMulti/src/peopleGoUp/application"
	"apiMulti/src/peopleGoUp/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IPeopleGoUp, rabbitRepo domain.IPeopleGoUpRabbitqm) *gin.Engine {
	r := gin.Default()

	CreatePeopleGoUp := application.NewCreatepeopleGoUp(rabbitRepo, repo)
	createPeopleGoUpController := NewCreatePeopleGoUpController(CreatePeopleGoUp)

	viewPeopleGoUp := application.NewViewPeopleGoUp(repo)
	viewPeopleGoUpController := NewViewPeopleGoUpController(viewPeopleGoUp)

	r.POST("/peopleGoUp", createPeopleGoUpController.Execute)
	r.GET("/peopleGoUpTest", viewPeopleGoUpController.Execute)

	return r
}
