package app

import (
	"log"
	"os"
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
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Router *gin.Engine
	DB     *mongo.Database // ajusta según tu tipo de DB
	Client *mongo.Client   // cliente Mongo
}

func NewApp() {
	config.ConfiguracionLog()
	defer config.CerrarLog()

	// Cargar .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	urlMongo := os.Getenv("URL_MONGO")

	db, cliente, err := config.ConnectMongo(urlMongo, "presto")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	router.SetTrustedProxies([]string{"127.0.0.1"})
	api := router.Group("api")
	initCliente(api, db)
	initTarifa(api, db)
	initMedidor(api, db)
	initLectura(api, db)
	initPago(api, db, cliente)
	initUsuario(api, db)
}

func (app *App) Run() {
	app.Router.Run(":5000")
}

func initCliente(api *gin.RouterGroup, db *mongo.Database) {
	repo := clienteRepository.NewClienteRepository(db)
	service := clienteService.NewClienteService(repo)
	controller := clienteController.NewClienteController(service)
	clienteRouter.ClienteRouter(api, controller)
}

func initTarifa(api *gin.RouterGroup, db *mongo.Database) {
	rangoRepo := tarifaRepository.NewRangoRepository(db)
	tarifaRepo := tarifaRepository.NewTarifaRepository(db)
	service := tarifaService.NewTarifaService(rangoRepo, tarifaRepo)
	controller := tarifaController.NewTarifaController(service)
	tarifaRouter.TarifaRouter(api, controller)
}

func initMedidor(api *gin.RouterGroup, db *mongo.Database) {
	repo := medidorRepository.NewMedidorRespository(db)
	service := medidorService.NewMedidoService(repo)
	controller := medidorController.NewMedidorController(service)
	medidorRouter.MedidorRouter(api, controller)
}

func initLectura(api *gin.RouterGroup, db *mongo.Database) {
	repoLectura := lecturaRepository.NewLecturaRepository(db)
	repoMedidor := medidorRepository.NewMedidorRespository(db)
	rangoRepo := tarifaRepository.NewRangoRepository(db)
	service := lecturaService.NewLecturaService(repoLectura, rangoRepo, repoMedidor)
	controller := lecturaController.NewLecturaController(service)
	lecturaRouter.LecturaRouter(api, controller)
}

func initPago(api *gin.RouterGroup, db *mongo.Database, cliente *mongo.Client) {
	pagoRepo := pagosRepository.NewPagoRepository(db)
	detallePagoRepo := pagosRepository.NewDetallePagoRepository(db)
	repoLectura := lecturaRepository.NewLecturaRepository(db)
	repoMedidor := medidorRepository.NewMedidorRespository(db)
	service := pagoService.NewPagoService(pagoRepo, repoLectura, repoMedidor, detallePagoRepo, cliente)
	controller := pagoController.NewPagoController(service)
	pagoRouter.PagoRouter(api, controller)
}

func initUsuario(api *gin.RouterGroup, db *mongo.Database) {
	repo := usuarioRepository.NewUsuarioRepository(db)
	service := usuarioService.NewUsuarioService(repo)
	controller := usuarioController.NewUsuarioController(service)
	usuarioRouter.UsuarioRouter(api, controller)
}
