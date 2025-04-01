package infraestructure

import (
	"apiMulti/src/transit/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTransitController struct {
	useCase *application.CreateTransit
}

func NewCreateTransitController(useCase *application.CreateTransit) *CreateTransitController {
	return &CreateTransitController{useCase: useCase}
}

type RequestBody struct {
	Cantidad int32 `json:"cantidad"`
	Tiempo   string `json:"tiempo"`
	Fecha    string `json:"fecha"`
}

func (ct_c *CreateTransitController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	err := ct_c.useCase.Execute(body.Cantidad, body.Tiempo, body.Fecha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar ek transito de personas", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transito de personas agregado correctamente"})
}