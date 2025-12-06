package enum

type FlagE string

const (
	FlagNuevo     FlagE = "nuevo"
	FlagEliminado FlagE = "elimiando"
)

type RolE string

const (
	RolAdministrador RolE = "ADMINISTRADOR"
	RolLecturador    RolE = "LECTURADOR"
)

type EstadoMedidor string

const (
	MedidorActivo        EstadoMedidor = "ACTIVO"        // Medidor funcionando normalmente
	MedidorInactivo      EstadoMedidor = "INACTIVO"      // Medidor retirado o sin uso
	MedidorMantenimiento EstadoMedidor = "MANTENIMIENTO" // Medidor temporalmente fuera de servicio
	MedidorEnCorte       EstadoMedidor = "EN_CORTE"      // Medidor cortado por falta de pago
)

type EstadoLectura string

const (
	LecturaPendiente EstadoLectura = "PENDIENTE" // Medidor funcionando normalmente
	LecturaPagado    EstadoLectura = "PAGADO"    // Medidor retirado o sin uso

)

type TipoPago string

const (
	TipoPagoEfectivo TipoPago = "EFECTIVO" // Medidor funcionando normalmente
	// Medidor retirado o sin uso

)
