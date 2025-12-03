package main

import (
	"context"
	"fmt"
	"prestoBackend/src/core/config"
	clienteController "prestoBackend/src/module/cliente/controller"
	clienteRepository "prestoBackend/src/module/cliente/repository"
	clienteRouter "prestoBackend/src/module/cliente/router"

	clienteService "prestoBackend/src/module/cliente/service"

	lecturaController "prestoBackend/src/module/lectura/controller"
	lecturaRepository "prestoBackend/src/module/lectura/repository"
	lecturaRouter "prestoBackend/src/module/lectura/router"
	lecturaService "prestoBackend/src/module/lectura/service"

	medidorController "prestoBackend/src/module/medidor/controller"
	medidorRepository "prestoBackend/src/module/medidor/repository"
	medidorRouter "prestoBackend/src/module/medidor/router"
	medidorService "prestoBackend/src/module/medidor/service"

	tarifaController "prestoBackend/src/module/tarifa/controller"
	tarifaRepository "prestoBackend/src/module/tarifa/repository"
	tarifaRouter "prestoBackend/src/module/tarifa/router"
	tarifaService "prestoBackend/src/module/tarifa/service"

	"github.com/gin-gonic/gin"
)

func main() {
	var url string = "mongodb://kanna:kanna@localhost:27017/presto?authSource=admin"
	db, cliente, err := config.ConnectMongo(url, "presto")

	if err != nil {
		fmt.Println(err)
	}
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
	lecturaService := lecturaService.NewLecturaService(lecturaRepository)
	lecturaController := lecturaController.NewLecturaController(lecturaService)
	lecturaRouter.LecturaRouter(api, lecturaController)

	//lectura
	medidorRepository := medidorRepository.NewMedidorRespository(db)
	medidorService := medidorService.NewMedidoService(medidorRepository)
	medidorController := medidorController.NewMedidorController(medidorService)
	medidorRouter.MedidorRouter(api, medidorController)

	//tarifa
	tarifaRepo := tarifaRepository.NewTarifaRepository(db)
	ranfoRepo := tarifaRepository.NewRangoRepository(db)
	tarifaService := tarifaService.NewTarifaService(ranfoRepo, tarifaRepo)
	tarifaController := tarifaController.NewTarifaController(tarifaService)
	tarifaRouter.TarifaRouter(api, tarifaController)

	router.Run(":5000")
}
