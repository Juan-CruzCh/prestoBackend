package main

import (
	"context"
	"prestoBackend/src/core/config"
	clienteController "prestoBackend/src/module/cliente/controller"
	clienteRepository "prestoBackend/src/module/cliente/repository"
	clienteRouter "prestoBackend/src/module/cliente/router"
	clienteService "prestoBackend/src/module/cliente/service"
	lecturaController "prestoBackend/src/module/lectura/controller"
	lecturaRepository "prestoBackend/src/module/lectura/repository"
	lecturaRouter "prestoBackend/src/module/lectura/router"
	lecturaService "prestoBackend/src/module/lectura/service"

	"github.com/gin-gonic/gin"
)

func main() {
	var url string = "mongodb://localhost:27017"
	db, cliente, _ := config.ConnectMongo(url, "presto")
	defer cliente.Disconnect(context.TODO())

	router := gin.Default()
	api := router.Group("api")

	//cliente
	clienteRepository := clienteRepository.NewClienteRepository(db)
	clienteService := clienteService.NewClienteService(clienteRepository)
	clienteController := clienteController.NewClienteController(clienteService)
	clienteRouter.ClienteRouter(api, clienteController)

	//lectura
	lecturaRepository := lecturaRepository.NewLecturaRepository(db)
	lecturaService := lecturaService.NewLecturaService(&lecturaRepository)
	lecturaController := lecturaController.NewLecturaController(lecturaService)
	lecturaRouter.LecturaRouter(api, lecturaController)

	router.Run()
}
