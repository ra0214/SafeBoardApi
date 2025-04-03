package infraestructure

import (
	"apiMulti/src/peopleGoUp/application"
	"log"
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
		log.Printf("Error en ViewPeopleGoUpController: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"details": "Error al obtener los datos",
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
