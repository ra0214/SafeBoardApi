package main

import (
	"github.com/gin-gonic/gin"

	"apiMulti/src/config"
	"apiMulti/src/config/middleware"
	"apiMulti/src/movement/infraestructure"
	goupInfra "apiMulti/src/peopleGoUp/infraestructure"
	godownInfra "apiMulti/src/peopleGoDown/infraestructure"
	userInfra "apiMulti/src/users/infraestructure"
	"log"
)

func main() {
	r := gin.Default()

	r.Use(middleware.NewCorsMiddleware())

	mysqlRepo := infraestructure.NewMySQL()
	userRepo := userInfra.NewMySQL()
	goupRepo := goupInfra.NewMySQL()
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

	godownRouter := godownInfra.SetupRouter(godownRepo, godownRabbitRepo)
	for _, route := range godownRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	r.SetTrustedProxies([]string{"127.0.0.1"})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
