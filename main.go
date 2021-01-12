package main

import (
	"log"

	"github.com/SanabriaDDi/api-redtwee/bd"
	"github.com/SanabriaDDi/api-redtwee/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
