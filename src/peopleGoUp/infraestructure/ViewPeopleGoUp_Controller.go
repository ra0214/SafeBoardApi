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

func (et_c *ViewPeopleGoUpController) Execute(c *gin.Context) {
	citas, err := et_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, citas)
}