package config

import (
	"log"
	"os"
)

var LogFile *os.File

func ConfiguracionLog() {
	var err error
	LogFile, err = os.OpenFile("presto.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error al abrir el archivo de log: %v", err)
	}
	log.SetOutput(LogFile)
}

func CerrarLog() {
	if LogFile != nil {
		LogFile.Close()
	}
}
