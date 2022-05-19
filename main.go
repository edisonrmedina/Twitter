package main

import (
	"log"

	"github.com/edisonrmedina/Twitter/bd"
	"github.com/edisonrmedina/Twitter/handlers"
)

func main() {
	if bd.CheckConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()

}
