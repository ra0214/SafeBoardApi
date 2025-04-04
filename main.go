package main

import (
	"github.com/gin-gonic/gin"

	"apiMulti/src/config"
	"apiMulti/src/config/websocket"
	"apiMulti/src/movement/infraestructure"
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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Inicializar repositorios MySQL
	userRepo := userInfra.NewMySQL()
	goupRepo := goupInfra.NewMySQL()
	godownRepo := godownInfra.NewMySQL()
	movementRepo := infraestructure.NewMySQL()

	// Inicializar RabbitMQ
	rabbitMQRepo, err := config.GetChannel()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer rabbitMQRepo.Close()

	// Inicializar repositorios RabbitMQ
	goupRabbitRepo := goupInfra.NewRabbitRepository(rabbitMQRepo.Ch)
	godownRabbitRepo := godownInfra.NewRabbitRepository(rabbitMQRepo.Ch)
	movementRabbitRepo := infraestructure.NewRabbitRepository(rabbitMQRepo.Ch)

	// Configurar rutas - Cada ruta se registra una sola vez
	// Users
	userRouter := userInfra.SetupRouter(userRepo)
	for _, route := range userRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	// PeopleGoUp
	goupRouter := goupInfra.SetupRouter(goupRepo, goupRabbitRepo)
	for _, route := range goupRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	// PeopleGoDown
	godownRouter := godownInfra.SetupRouter(godownRepo, godownRabbitRepo)
	for _, route := range godownRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	// Movement
	movementRouter := infraestructure.SetupRouter(movementRepo, movementRabbitRepo)
	for _, route := range movementRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	// WebSocket
	r.GET("/ws", websocket.HandleWebSocket)

	// Configurar servidor
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Iniciar servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
