package infraestructure

import (
	"apiMulti/src/peopleGoDown/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePeopleGoDownController struct {
	createPeopleGoDown *application.CreatePeopleGoDown
}

func NewCreatePeopleGoDownController(useCase *application.CreatePeopleGoDown) *CreatePeopleGoDownController {
	return &CreatePeopleGoDownController{
		createPeopleGoDown: useCase,
	}
}

func (cc *CreatePeopleGoDownController) Execute(c *gin.Context) {
	var request struct {
		ESP32ID  string `json:"esp32_id"`
		Cantidad int32  `json:"conteo"`
	}

	if err := c.BindJSON(&request); err != nil {
		log.Printf("[Controller] Error al parsear request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.createPeopleGoDown.Execute(request.ESP32ID, request.Cantidad)
	if err != nil {
		log.Printf("[Controller] Error en Execute: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registro creado exitosamente"})
}
