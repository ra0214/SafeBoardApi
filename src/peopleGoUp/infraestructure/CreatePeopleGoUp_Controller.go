package infraestructure

import (
	"apiMulti/src/peopleGoUp/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePeopleGoUpController struct {
	useCase *application.CreatepeopleGoUp
}

func NewCreatePeopleGoUpController(useCase *application.CreatepeopleGoUp) *CreatePeopleGoUpController {
	return &CreatePeopleGoUpController{useCase: useCase}
}

type RequestBody struct {
	Cantidad int32 `json:"cantidad"`
}

func (ct_c *CreatePeopleGoUpController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	err := ct_c.useCase.Execute(body.Cantidad)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el transito de personas", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Cuenta de personas agregada correctamente"})
}