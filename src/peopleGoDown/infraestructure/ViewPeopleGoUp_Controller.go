package infraestructure

import (
	"apiMulti/src/peopleGoDown/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewPeopleGoDownController struct {
	useCase *application.ViewPeopleGoDown
}

func NewViewPeopleGoDownController(useCase *application.ViewPeopleGoDown) *ViewPeopleGoDownController {
	return &ViewPeopleGoDownController{useCase: useCase}
}

func (et_c *ViewPeopleGoDownController) Execute(c *gin.Context) {
	citas, err := et_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, citas)
}