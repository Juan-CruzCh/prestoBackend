package src

import (
	"fmt"
	"log"
	"net/http"
	"prestoBackend/src/core/config"
	"prestoBackend/src/core/middleware"
	"prestoBackend/src/core/utils"
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

	autenticacionController "prestoBackend/src/module/autenticacion/controller"
	autenticacionRouter "prestoBackend/src/module/autenticacion/router"
	autenticacionService "prestoBackend/src/module/autenticacion/service"

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
	ServerMux    *http.ServeMux
	DB           *mongo.Database
	Client       *mongo.Client
	Repositories *Repositories
}

func NewApp(urlMongo string) *App {

	db, cliente, err := config.ConnectMongo(urlMongo, "presto")
	if err != nil {
		log.Fatal(err)
	}

	/*router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))*/
	serverMux := http.NewServeMux()

	app := &App{
		ServerMux:    serverMux,
		DB:           db,
		Client:       cliente,
		Repositories: NewRepositories(db),
	}
	initCliente(app)
	initTarifa(app)
	initMedidor(app)
	initLectura(app)
	initPago(app)
	initUsuario(app)
	initAutenticacion(app)
	return app
}

func (app *App) Run(port string) {
	log.Printf("Servidor corriendo en http://localhost:%s", port)
	fmt.Println("Servidor corriendo en http://localhost:%s", port)

	handlerConMiddleware := utils.LoggingMiddleware(app.ServerMux)
	middlewareCors := middleware.EnableCORS(handlerConMiddleware)
	err := http.ListenAndServe(":"+port, middlewareCors)
	if err != nil {
		log.Fatal(err)
	}
}

func initCliente(app *App) {
	service := clienteService.NewClienteService(app.Repositories.ClienteRepository, app.Repositories.MedidorRepository)
	controller := clienteController.NewClienteController(service)
	clienteRouter.NewClienteRouter(app.ServerMux, controller)

}

func initTarifa(app *App) {
	service := tarifaService.NewTarifaService(app.Repositories.RangoRepository, app.Repositories.TarifaRepository)
	controller := tarifaController.NewTarifaController(service)
	r := tarifaRouter.NewTarifaRouter(app.ServerMux, controller)
	r.TarifaRouter()
}

func initMedidor(app *App) {
	service := medidorService.NewMedidoService(app.Repositories.MedidorRepository)
	controller := medidorController.NewMedidorController(service)
	r := medidorRouter.NewMedidorRouter(app.ServerMux, controller)
	r.MedidorRouter()
}

func initLectura(app *App) {
	service := lecturaService.NewLecturaService(app.Repositories.LecturaRepository, app.Repositories.RangoRepository, app.Repositories.MedidorRepository)
	controller := lecturaController.NewLecturaController(service)
	r := lecturaRouter.NewLecturaRouter(app.ServerMux, controller)
	r.LecturaRouter()
}

func initPago(app *App) {
	service := pagoService.NewPagoService(app.Repositories.PagoRepository, app.Repositories.LecturaRepository, app.Repositories.MedidorRepository, app.Repositories.DetallePagoRepository)
	controller := pagoController.NewPagoController(service)
	r := pagoRouter.NewPagoRouter(app.ServerMux, controller)
	r.PagoRouter()
}

func initUsuario(app *App) {
	service := usuarioService.NewUsuarioService(app.Repositories.UsuarioRepository)
	controller := usuarioController.NewUsuarioController(service)
	r := usuarioRouter.NewUsuarioRouter(app.ServerMux, controller)
	r.UsuarioRouter()
}

func initAutenticacion(app *App) {
	service := autenticacionService.NewAutenticacionService(app.Repositories.UsuarioRepository)
	controller := autenticacionController.NewAutenticacionController(service)
	r := autenticacionRouter.NewAutenticacionRouter(app.ServerMux, controller)
	r.AutenticacionRouter()
}
