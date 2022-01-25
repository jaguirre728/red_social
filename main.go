package main

import (
	"log"

	"github.com/jaguirre728/red_social.git/bd"
	"github.com/jaguirre728/red_social.git/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Manejadores()
}
