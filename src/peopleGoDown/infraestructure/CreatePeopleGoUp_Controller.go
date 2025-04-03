package infraestructure

import (
	"apiMulti/src/peopleGoDown/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePeopleGoDownController struct {
	useCase *application.CreatepeopleGoDown
}

func NewCreatePeopleGoDownController(useCase *application.CreatepeopleGoDown) *CreatePeopleGoDownController {
	return &CreatePeopleGoDownController{useCase: useCase}
}

type RequestBody struct {
	Esp32ID string `json:"esp32_id"`
	Conteo int32 `json:"conteo"`
}

func (ct_c *CreatePeopleGoDownController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	err := ct_c.useCase.Execute(body.Esp32ID ,body.Conteo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el transito de personas", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Cuenta de personas agregada correctamente"})
}