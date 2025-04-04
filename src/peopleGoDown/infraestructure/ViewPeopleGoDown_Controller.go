package infraestructure

import (
	"apiMulti/src/peopleGoDown/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewPeopleGoDownController struct {
	viewPeopleGoDown *application.ViewPeopleGoDown
}

func NewViewPeopleGoDownController(useCase *application.ViewPeopleGoDown) *ViewPeopleGoDownController {
	return &ViewPeopleGoDownController{
		viewPeopleGoDown: useCase,
	}
}

func (vc *ViewPeopleGoDownController) Execute(c *gin.Context) {
	log.Printf("[Controller] Iniciando GetAll de PeopleGoDown")

	data, err := vc.viewPeopleGoDown.Execute()
	if err != nil {
		log.Printf("[Controller] Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[Controller] Datos obtenidos exitosamente: %d registros", len(data))
	c.JSON(http.StatusOK, data)
}
