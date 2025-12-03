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
