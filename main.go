package main

import (
	"log"

	"github.com/alaanmdq1/gobloging/bd"
	"github.com/alaanmdq1/gobloging/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
