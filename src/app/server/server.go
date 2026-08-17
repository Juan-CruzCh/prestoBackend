package server

import (
	"fmt"
	"log"
	"net/http"
	"prestoBackend/src/app/config"
	"prestoBackend/src/app/database"
	"prestoBackend/src/app/middleware"
	clienteController "prestoBackend/src/internal/cliente/controller"
	clienteRepository "prestoBackend/src/internal/cliente/repository"
	clienteRouter "prestoBackend/src/internal/cliente/router"

	clienteService "prestoBackend/src/internal/cliente/service"

	lecturaController "prestoBackend/src/internal/lectura/controller"
	lecturaRepository "prestoBackend/src/internal/lectura/repository"
	lecturaRouter "prestoBackend/src/internal/lectura/router"
	lecturaService "prestoBackend/src/internal/lectura/service"

	medidorController "prestoBackend/src/internal/medidor/controller"
	medidorRepository "prestoBackend/src/internal/medidor/repository"
	medidorRouter "prestoBackend/src/internal/medidor/router"
	medidorService "prestoBackend/src/internal/medidor/service"

	tarifaController "prestoBackend/src/internal/tarifa/controller"
	tarifaRepository "prestoBackend/src/internal/tarifa/repository"
	tarifaRouter "prestoBackend/src/internal/tarifa/router"
	tarifaService "prestoBackend/src/internal/tarifa/service"

	pagoController "prestoBackend/src/internal/pago/controller"
	pagosRepository "prestoBackend/src/internal/pago/repository"

	pagoRouter "prestoBackend/src/internal/pago/router"
	pagoService "prestoBackend/src/internal/pago/service"

	usuarioController "prestoBackend/src/internal/usuario/controller"
	usuarioRepository "prestoBackend/src/internal/usuario/repository"
	usuarioRouter "prestoBackend/src/internal/usuario/router"
	usuarioService "prestoBackend/src/internal/usuario/service"

	autenticacionController "prestoBackend/src/internal/autenticacion/controller"
	autenticacionRouter "prestoBackend/src/internal/autenticacion/router"
	autenticacionService "prestoBackend/src/internal/autenticacion/service"

	cajaController "prestoBackend/src/internal/caja/controller"
	cajaRepository "prestoBackend/src/internal/caja/repository"
	cajacaRouter "prestoBackend/src/internal/caja/router"
	cajaService "prestoBackend/src/internal/caja/service"

	cajaChicaController "prestoBackend/src/internal/cajaChica/controller"
	cajaChicaRepository "prestoBackend/src/internal/cajaChica/repository"
	cajaChicaRouter "prestoBackend/src/internal/cajaChica/router"
	cajaChicaService "prestoBackend/src/internal/cajaChica/service"

	gastoController "prestoBackend/src/internal/gasto/controller"
	gastoRepository "prestoBackend/src/internal/gasto/repository"
	gastoRouter "prestoBackend/src/internal/gasto/router"
	gastoService "prestoBackend/src/internal/gasto/service"
	servicioController "prestoBackend/src/internal/servicio/controller"
	servicoRepository "prestoBackend/src/internal/servicio/repository"
	servicioRouter "prestoBackend/src/internal/servicio/router"
	servicioService "prestoBackend/src/internal/servicio/service"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repositories struct {
	ClienteRepository        clienteRepository.ClienteRepository
	MedidorRepository        medidorRepository.MedidorRepository
	LecturaRepository        lecturaRepository.LecturaRepository
	TarifaRepository         tarifaRepository.TarifaRepository
	RangoRepository          tarifaRepository.RangoRepository
	PagoRepository           pagosRepository.PagoRepository
	DetallePagoRepository    pagosRepository.DetallePagoRepository
	UsuarioRepository        usuarioRepository.UsuarioRepository
	CajaRepository           cajaRepository.Caja
	CajaChicaRepository      cajaChicaRepository.CajaChica
	CategoriaGastoRepository gastoRepository.CategoriaGasto
	GastoRepository          gastoRepository.Gasto
	ServicoRepository        servicoRepository.Servicio
	DetalleServicoRepository servicoRepository.DetalleServicio
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		ClienteRepository:        clienteRepository.NewClienteRepository(db),
		MedidorRepository:        medidorRepository.NewMedidorRespository(db),
		LecturaRepository:        lecturaRepository.NewLecturaRepository(db),
		TarifaRepository:         tarifaRepository.NewTarifaRepository(db),
		RangoRepository:          tarifaRepository.NewRangoRepository(db),
		PagoRepository:           pagosRepository.NewPagoRepository(db),
		UsuarioRepository:        usuarioRepository.NewUsuarioRepository(db),
		DetallePagoRepository:    pagosRepository.NewDetallePagoRepository(db),
		CajaRepository:           cajaRepository.NewCajaRepository(db),
		CajaChicaRepository:      cajaChicaRepository.NewCajaChicaRepository(db),
		CategoriaGastoRepository: gastoRepository.NewCategoriaGastoRepository(db),
		GastoRepository:          gastoRepository.NewGastoRepository(db),
		ServicoRepository:        servicoRepository.NewServicioRepository(db),
		DetalleServicoRepository: servicoRepository.NewDetalleServicioRepository(db),
	}
}

type App struct {
	ServerMux    *http.ServeMux
	Repositories *Repositories
	Validate     *validator.Validate
	Cliente      *mongo.Client
}

func NewApp() *App {

	db, cliente, err := database.Connection(config.AppEnv.MongoURI, config.AppEnv.Database)
	if err != nil {
		log.Fatal(err)
	}
	validate := validator.New()
	serverMux := http.NewServeMux()

	app := &App{
		ServerMux:    serverMux,
		Repositories: NewRepositories(db),
		Validate:     validate,
		Cliente:      cliente,
	}
	initCliente(app)
	initTarifa(app)
	initMedidor(app)
	initLectura(app)
	initPago(app)
	initUsuario(app)
	initAutenticacion(app)
	initCajaChica(app)
	initCaja(app)
	initServicio(app)
	initGasto(app)
	return app
}

func (app *App) Run() {
	log.Printf("Servidor corriendo en http://localhost:%s", config.AppEnv.Port)
	fmt.Println("Servidor corriendo en http://localhost:%s", config.AppEnv.Port)
	var handler http.Handler = app.ServerMux
	handler = middleware.VerificarAutenticacion(app.Repositories.UsuarioRepository)(handler)
	handler = middleware.LoggingMiddleware(handler)
	handler = middleware.EnableCORS(handler)
	err := http.ListenAndServe(":"+config.AppEnv.Port, handler)
	if err != nil {
		log.Fatal(err)
	}
}

func initCliente(app *App) {
	service := clienteService.NewClienteService(app.Repositories.ClienteRepository, app.Repositories.MedidorRepository)
	controller := clienteController.NewClienteController(service, app.Validate)
	clienteRouter.NewClienteRouter(app.ServerMux, controller)
}

func initTarifa(app *App) {
	service := tarifaService.NewTarifaService(app.Repositories.RangoRepository, app.Repositories.TarifaRepository)
	controller := tarifaController.NewTarifaController(service, app.Validate)
	r := tarifaRouter.NewTarifaRouter(app.ServerMux, controller)
	r.TarifaRouter()
}

func initMedidor(app *App) {
	service := medidorService.NewMedidoService(app.Repositories.MedidorRepository)
	controller := medidorController.NewMedidorController(service, app.Validate)
	r := medidorRouter.NewMedidorRouter(app.ServerMux, controller)
	r.MedidorRouter()
}

func initLectura(app *App) {
	service := lecturaService.NewLecturaService(app.Repositories.LecturaRepository, app.Repositories.RangoRepository, app.Repositories.MedidorRepository, app.Cliente)
	controller := lecturaController.NewLecturaController(service, app.Validate)
	r := lecturaRouter.NewLecturaRouter(app.ServerMux, controller)
	r.LecturaRouter()
}

func initPago(app *App) {
	service := pagoService.NewPagoService(app.Repositories.PagoRepository, app.Repositories.LecturaRepository, app.Repositories.MedidorRepository, app.Repositories.DetallePagoRepository, app.Repositories.CajaRepository, app.Cliente)
	controller := pagoController.NewPagoController(service, app.Validate)
	r := pagoRouter.NewPagoRouter(app.ServerMux, controller)
	r.PagoRouter()
}

func initUsuario(app *App) {
	service := usuarioService.NewUsuarioService(app.Repositories.UsuarioRepository)
	controller := usuarioController.NewUsuarioController(service, app.Validate)
	r := usuarioRouter.NewUsuarioRouter(app.ServerMux, controller)
	r.UsuarioRouter()
}

func initAutenticacion(app *App) {
	service := autenticacionService.NewAutenticacionService(app.Repositories.UsuarioRepository)
	controller := autenticacionController.NewAutenticacionController(service, app.Validate)
	r := autenticacionRouter.NewAutenticacionRouter(app.ServerMux, controller)
	r.AutenticacionRouter()
}

func initCaja(app *App) {
	service := cajaService.NewCajaService(app.Repositories.CajaRepository, app.Repositories.PagoRepository, app.Cliente)
	controller := cajaController.NewCajaController(service, app.Validate)
	cajacaRouter.NewCajaRouter(app.ServerMux, controller)
}

func initServicio(app *App) {
	service := servicioService.NewServicioService(app.Repositories.ServicoRepository, app.Cliente)
	controller := servicioController.NewServicioController(service, app.Validate)
	servicioRouter.NewServicioRouter(app.ServerMux, controller)
}

func initCajaChica(app *App) {
	service := cajaChicaService.NewCajaChicaService(app.Repositories.CajaChicaRepository, app.Cliente)
	controller := cajaChicaController.NewCajaChicaController(service, app.Validate)
	cajaChicaRouter.NewCajaChicaRouter(app.ServerMux, controller)
}

func initGasto(app *App) {
	service := gastoService.NewGastoService(app.Repositories.GastoRepository, app.Repositories.CajaChicaRepository, app.Cliente)
	controller := gastoController.NewGastoController(service, app.Validate)
	gastoRouter.NewGastoRouter(app.ServerMux, controller)

	serviceCategoria := gastoService.NewCategoriaGastoService(app.Repositories.CategoriaGastoRepository, app.Cliente)
	controllerCategoria := gastoController.NewCategoriaGastoController(serviceCategoria, app.Validate)
	gastoRouter.NewCategoriaGastoRouter(app.ServerMux, controllerCategoria)

}
