package routes

import (
	db "Persona-Mysql_Rest/DB"
	"Persona-Mysql_Rest/controller"
	"Persona-Mysql_Rest/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Miki")
}

func RestPersonas(w http.ResponseWriter, r *http.Request) {
	ListaPersonas := model.AllPersonas{}
	conn, err := db.ConexionBD()
	if err != nil {
		panic(err)

	}

	ListaPersonas, err = controller.GetPersonas1(r.Context(), conn.SQL)
	if err != nil {
		panic(err)

	}

	//Envia o Responde de Tipo Json
	w.Header().Set("Content-Type", "application/json")
	//Enviamos un Codigo de estado correcto
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&ListaPersonas)
}

func RestSavePersonas(w http.ResponseWriter, r *http.Request) {

	var nuevaPersona model.Persona
	//Permite manejar las entradas y salidad del Servidor
	datos, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte Datos Validos")
	}
	/*
		Si los datos son validos
		se asigna a la snuevaPersona los datos recibidos
		para poder manipular .
	*/
	json.Unmarshal(datos, &nuevaPersona)

	/*
		Agregamos el id de manera manual,
		Se basa en el tamano del arreglo para asignar el idPersona al nuevo dato que se agrega
	*/
	conn, _ := db.ConexionBD()
	err = controller.CreatePerson(r.Context(), conn.SQL, &nuevaPersona)
	if err != nil {
		panic(err)

	}
	//Enviamos la nueva Persona que se agrego al arreglo en formato
	w.Header().Set("Content-Type", "application/json")
	//Enviamos un Codigo de estado correcto
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(nuevaPersona)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dato, err := strconv.Atoi(vars["IdPersona"])

	if err != nil {
		fmt.Fprintf(w, "ID Invalido")
		return
	}
	conn, _ := db.ConexionBD()

	controller.EliminarPersona1(r.Context(), conn.SQL, int64(dato))

	//_, err := p.repo.Delete(r.Context(), int64(i))
	fmt.Fprintf(w, "La Persona con ID %v fue eliminada", dato)

}

func BuscarPersonaId(w http.ResponseWriter, r *http.Request) {
	//extrae los valores del url
	vars := mux.Vars(r)
	//Convierte un String a un Entero
	//Puede o no ocurrir un error
	dato, err := strconv.Atoi(vars["IdPersona"])
	if err != nil {
		fmt.Fprintf(w, "ID Invalido")
		return
	}

	conn, _ := db.ConexionBD()

	per := controller.BuscarPersona1(r.Context(), conn.SQL, int64(dato))

	//_, err := p.repo.Delete(r.Context(), int64(i))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(per)
}

func UpdatePersonaID(w http.ResponseWriter, r *http.Request) {
	var actualizarPersona model.Persona

	datos, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte Datos Validos")
	}
	/*
		Si los datos son validos
		se asigna a la snuevaPersona los datos recibidos
		para poder manipular .
	*/
	json.Unmarshal(datos, &actualizarPersona)
	//asignar los datos del body a la variable actualizarPersona para manipularlso
	conn, _ := db.ConexionBD()
	err = controller.UpdatePersona1(r.Context(), conn.SQL, &actualizarPersona)
	if err != nil {
		panic(err)

	}
	//Enviamos la nueva Persona que se agrego al arreglo en formato
	w.Header().Set("Content-Type", "application/json")
	//Enviamos un Codigo de estado correcto
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Coreecto")

}
