package main

import (
	route "Persona-Mysql_Rest/Routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// var nuevaPersona model.Persona

func main() {
	fmt.Println("hola mundo")

	//EL url debe tener una ruta valida
	router := mux.NewRouter().StrictSlash(true)

	//ruta inicial
	router.HandleFunc("/", route.Index1)

	//Responder a una ruta y por un get la lista de Personas
	router.HandleFunc("/listaPersona", route.RestPersonas).Methods("GET")

	//Ruta para Agregar una nuevaPersona
	router.HandleFunc("/crearPersona", route.RestSavePersonas).Methods("POST")

	//Buscar una Persona por su Id
	router.HandleFunc("/buscarPersona/{IdPersona}", route.BuscarPersonaId).Methods("GET")

	//Buscar una Persona por su Id
	router.HandleFunc("/eliminarPersona/{IdPersona}", route.Delete).Methods("DELETE")

	//Buscar una Persona por su Id
	router.HandleFunc("/actualizarPersona", route.UpdatePersonaID).Methods("PUT")

	//iniciar el Servidor en el  puerto 8080
	servidor := http.ListenAndServe(":8080", router)
	//Log
	log.Fatal(servidor)

}
