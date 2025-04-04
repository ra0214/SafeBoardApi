package main

import (
	"github.com/gin-gonic/gin"

	"apiMulti/src/config"
	"apiMulti/src/movement/infraestructure"
	"apiMulti/src/peopleGoDown/application"
	godownInfra "apiMulti/src/peopleGoDown/infraestructure"
	goupInfra "apiMulti/src/peopleGoUp/infraestructure"
	userInfra "apiMulti/src/users/infraestructure"
	"log"
)

func main() {
	r := gin.Default()

	// Configuración CORS global
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Inicializar repositorios
	userRepo := userInfra.NewMySQL()
	goupRepo := goupInfra.NewMySQL()
	mysqlRepo := infraestructure.NewMySQL()
	godownRepo := godownInfra.NewMySQL()

	rabbitMQRepo, err := config.GetChannel()
	if err != nil {
		log.Fatal("Error al conectar con RabbitMQ: %v", err)
	}
	defer rabbitMQRepo.Close()

	rabbitRepo := infraestructure.NewRabbitRepository(rabbitMQRepo.Ch)

	transitRouter := infraestructure.SetupRouter(mysqlRepo, rabbitRepo)
	for _, route := range transitRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	userRouter := userInfra.SetupRouter(userRepo)
	for _, route := range userRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	goupRabbitRepo := goupInfra.NewRabbitRepository(rabbitMQRepo.Ch)

	goupRouter := goupInfra.SetupRouter(goupRepo, goupRabbitRepo)
	for _, route := range goupRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	godownRabbitRepo := godownInfra.NewRabbitRepository(rabbitMQRepo.Ch)

	// Crear los casos de uso
	createPeopleGoDown := application.NewCreatePeopleGoDown(godownRabbitRepo, godownRepo)
	viewPeopleGoDown := application.NewViewPeopleGoDown(godownRepo)

	// Crear los controladores
	createPeopleGoDownController := godownInfra.NewCreatePeopleGoDownController(createPeopleGoDown)
	viewPeopleGoDownController := godownInfra.NewViewPeopleGoDownController(viewPeopleGoDown)

	// Registrar las rutas
	r.POST("/peopleGoDown", createPeopleGoDownController.Execute)
	r.GET("/peopleGoDown", viewPeopleGoDownController.Execute)

	r.SetTrustedProxies([]string{"127.0.0.1"})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
