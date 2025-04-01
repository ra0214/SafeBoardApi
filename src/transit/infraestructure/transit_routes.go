package infraestructure

import (
	"apiMulti/src/transit/application"
	"apiMulti/src/transit/domain"
	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.ITransit, rabbitRepo domain.ITransitRabbitqm) *gin.Engine {
	r := gin.Default()

	CreateTransit := application.NewCreateTransit(rabbitRepo, repo)
	createTransitController := NewCreateTransitController(CreateTransit)

	viewTransit := application.NewViewTransit(repo)
	viewTransitController := NewViewTransitController(viewTransit)

	r.POST("/transit", createTransitController.Execute)
	r.GET("/transit", viewTransitController.Execute)

	return r
}