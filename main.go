package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"example/golang-rest-api/controller"
	"example/golang-rest-api/model"
)

var router *mux.Router

func initHandlers() {

	noteRouter := router.PathPrefix("/notes").Subrouter()
	noteRouter.HandleFunc("/", controller.NoteGetAllController).Methods(http.MethodGet)
	noteRouter.HandleFunc("/{tag}", controller.NoteGetByTagController).Methods(http.MethodGet)
	noteRouter.HandleFunc("/", controller.NoteCreateController).Methods(http.MethodPost)
}

func Start() {
	router = mux.NewRouter()
	
	initHandlers()
	fmt.Printf("router initialized and listening on 4000\n")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func main() {
	model.Init()
	Start()
}
