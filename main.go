package main

import (
	"context"
	"log"
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

	pagoController "prestoBackend/src/module/pago/controller"
	pagosRepository "prestoBackend/src/module/pago/repository"
	pagoRouter "prestoBackend/src/module/pago/router"
	pagoService "prestoBackend/src/module/pago/service"

	usuarioController "prestoBackend/src/module/usuario/controller"
	usuarioRepository "prestoBackend/src/module/usuario/repository"
	usuarioRouter "prestoBackend/src/module/usuario/router"
	usuarioService "prestoBackend/src/module/usuario/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConfiguracionLog()
	defer config.CerrarLog()

	var url string = "mongodb://kanna:kanna@localhost:27017/presto?authSource=admin"
	//var url string = "mongodb://localhost:27017"
	db, cliente, err := config.ConnectMongo(url, "presto")
	if err != nil {
		log.Println(err)
	}

	defer cliente.Disconnect(context.TODO())
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	router.SetTrustedProxies([]string{"127.0.0.1"})
	api := router.Group("api")

	//cliente
	clienteRepository := clienteRepository.NewClienteRepository(db)
	clienteService := clienteService.NewClienteService(clienteRepository)
	clienteController := clienteController.NewClienteController(clienteService)
	clienteRouter.ClienteRouter(api, clienteController)

	//tarifa
	tarifaRepo := tarifaRepository.NewTarifaRepository(db)
	rangoRepo := tarifaRepository.NewRangoRepository(db)
	tarifaService := tarifaService.NewTarifaService(rangoRepo, tarifaRepo)
	tarifaController := tarifaController.NewTarifaController(tarifaService)
	tarifaRouter.TarifaRouter(api, tarifaController)

	//lectura
	medidorRepository := medidorRepository.NewMedidorRespository(db)
	medidorService := medidorService.NewMedidoService(medidorRepository)
	medidorController := medidorController.NewMedidorController(medidorService)
	medidorRouter.MedidorRouter(api, medidorController)
	//lectura
	lecturaRepository := lecturaRepository.NewLecturaRepository(db)
	lecturaService := lecturaService.NewLecturaService(lecturaRepository, rangoRepo, medidorRepository)
	lecturaController := lecturaController.NewLecturaController(lecturaService)
	lecturaRouter.LecturaRouter(api, lecturaController)

	pagoRepository := pagosRepository.NewPagoRepository(db)
	detallepagoRepository := pagosRepository.NewDetallePagoRepository(db)
	pagoService := pagoService.NewPagoService(pagoRepository, lecturaRepository, medidorRepository, detallepagoRepository, cliente)
	pagoController := pagoController.NewPagoController(pagoService)
	pagoRouter.PagoRouter(api, pagoController)

	usuarioRepository := usuarioRepository.NewUsuarioRepository(db)

	usuarioService := usuarioService.NewUsuarioService(usuarioRepository)
	usuarioController := usuarioController.NewUsuarioController(usuarioService)
	usuarioRouter.UsuarioRouter(api, usuarioController)

	router.Run(":5000")
}
