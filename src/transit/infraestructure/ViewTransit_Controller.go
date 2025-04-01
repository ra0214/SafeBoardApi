package infraestructure

import (
	"apiMulti/src/transit/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewTransitController struct {
	useCase *application.ViewTransit
}

func NewViewTransitController(useCase *application.ViewTransit) *ViewTransitController {
	return &ViewTransitController{useCase: useCase}
}

func (et_c *ViewTransitController) Execute(c *gin.Context) {
	citas, err := et_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, citas)
}