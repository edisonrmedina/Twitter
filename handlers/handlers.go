package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	//por aqui manejadores las rutas

	router := mux.NewRouter()
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	//permisos para las direcciones ip
	handler := cors.AllowAll().Handler(router)
	//escucha el puerto, puerto y el handler
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
