package infraestructure

import (
	"apiMulti/src/peopleGoUp/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewPeopleGoUpController struct {
	useCase *application.ViewPeopleGoUp
}

func NewViewPeopleGoUpController(useCase *application.ViewPeopleGoUp) *ViewPeopleGoUpController {
	return &ViewPeopleGoUpController{useCase: useCase}
}

func (vc *ViewPeopleGoUpController) Execute(c *gin.Context) {
	data, err := vc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
