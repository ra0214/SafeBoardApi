package infraestructure

import (
	"apiMulti/src/movement/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMovementController struct {
	useCase *application.CreateMovement
}

func NewCreateMovementController(useCase *application.CreateMovement) *CreateMovementController {
	return &CreateMovementController{useCase: useCase}
}

type RequestBody struct {
	Esp32ID     string  `json:"esp32_id"`
	Aceleracion float64 `json:"aceleracion"`
}

func (ct_c *CreateMovementController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	err := ct_c.useCase.Execute(body.Esp32ID, body.Aceleracion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el transito de personas", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transito de personas agregado correctamente"})
}