package infraestructure

import (
	"apiMulti/src/movement/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewMovementController struct {
	useCase *application.ViewMovement
}

func NewViewMovementController(useCase *application.ViewMovement) *ViewMovementController {
	return &ViewMovementController{useCase: useCase}
}

func (et_c *ViewMovementController) Execute(c *gin.Context) {
	movements, err := et_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los movimientos", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movements)
}
