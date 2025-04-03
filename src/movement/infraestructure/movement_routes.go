package infraestructure

import (
	"apiMulti/src/movement/application"
	"apiMulti/src/movement/domain"
	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IMovement, rabbitRepo domain.IMovementRabbitqm) *gin.Engine {
	r := gin.Default()

	CreateMovement := application.NewCreateMovement(rabbitRepo, repo)
	createMovementController := NewCreateMovementController(CreateMovement)

	viewMovement := application.NewViewMovement(repo)
	viewMovementController := NewViewMovementController(viewMovement)

	r.POST("/movement", createMovementController.Execute)
	r.GET("/movement", viewMovementController.Execute)

	return r
}