package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/SanabriaDDi/api-redtwee/middlew"
	"github.com/SanabriaDDi/api-redtwee/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto y pongo a escuchar mi servidor*/
func Manejadores() {

	//Captura http
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//Setea permisos (a todo el mundo)
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
