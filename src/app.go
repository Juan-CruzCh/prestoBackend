package src

import (
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
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repositories struct {
	ClienteRepository     clienteRepository.ClienteRepository
	MedidorRepository     medidorRepository.MedidorRepository
	LecturaRepository     lecturaRepository.LecturaRepository
	TarifaRepository      tarifaRepository.TarifaRepository
	RangoRepository       tarifaRepository.RangoRepository
	PagoRepository        pagosRepository.PagoRepository
	DetallePagoRepository pagosRepository.DetallePagoRepository
	UsuarioRepository     usuarioRepository.UsuarioRepository
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		ClienteRepository:     clienteRepository.NewClienteRepository(db),
		MedidorRepository:     medidorRepository.NewMedidorRespository(db),
		LecturaRepository:     lecturaRepository.NewLecturaRepository(db),
		TarifaRepository:      tarifaRepository.NewTarifaRepository(db),
		RangoRepository:       tarifaRepository.NewRangoRepository(db),
		PagoRepository:        pagosRepository.NewPagoRepository(db),
		UsuarioRepository:     usuarioRepository.NewUsuarioRepository(db),
		DetallePagoRepository: pagosRepository.NewDetallePagoRepository(db),
	}
}

type App struct {
	Router       *gin.Engine
	DB           *mongo.Database
	Client       *mongo.Client
	Repositories *Repositories
}

func NewApp(urlMongo string) *App {

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

	app := &App{
		Router:       router,
		DB:           db,
		Client:       cliente,
		Repositories: NewRepositories(db),
	}
	initCliente(api, app)
	initTarifa(api, app)
	initMedidor(api, app)
	initLectura(api, app)
	initPago(api, app)
	initUsuario(api, app)

	return app
}

func (app *App) Run(port string) {
	app.Router.Run(":" + port)
}

func initCliente(api *gin.RouterGroup, app *App) {
	service := clienteService.NewClienteService(app.Repositories.ClienteRepository)
	controller := clienteController.NewClienteController(service)
	clienteRouter.ClienteRouter(api, controller)
}

func initTarifa(api *gin.RouterGroup, app *App) {

	service := tarifaService.NewTarifaService(app.Repositories.RangoRepository, app.Repositories.TarifaRepository)
	controller := tarifaController.NewTarifaController(service)
	tarifaRouter.TarifaRouter(api, controller)
}

func initMedidor(api *gin.RouterGroup, app *App) {

	service := medidorService.NewMedidoService(app.Repositories.MedidorRepository)
	controller := medidorController.NewMedidorController(service)
	medidorRouter.MedidorRouter(api, controller)
}

func initLectura(api *gin.RouterGroup, app *App) {

	service := lecturaService.NewLecturaService(app.Repositories.LecturaRepository, app.Repositories.RangoRepository, app.Repositories.MedidorRepository)
	controller := lecturaController.NewLecturaController(service)
	lecturaRouter.LecturaRouter(api, controller)
}

func initPago(api *gin.RouterGroup, app *App) {

	service := pagoService.NewPagoService(app.Repositories.PagoRepository, app.Repositories.LecturaRepository, app.Repositories.MedidorRepository, app.Repositories.DetallePagoRepository)
	controller := pagoController.NewPagoController(service)
	pagoRouter.PagoRouter(api, controller)
}

func initUsuario(api *gin.RouterGroup, app *App) {
	service := usuarioService.NewUsuarioService(app.Repositories.UsuarioRepository)
	controller := usuarioController.NewUsuarioController(service)
	usuarioRouter.UsuarioRouter(api, controller)
}
