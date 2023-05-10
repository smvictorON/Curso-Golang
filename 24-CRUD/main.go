package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	//mux é uma biblioteca robusta para usar com req http
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//Methods só diz o método que vai ser usado, "POST" escrito funciona tbm.
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.FindUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.FindUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Escutando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
