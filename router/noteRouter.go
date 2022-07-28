package router

import (
	"example/golang-rest-api/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func registerNoteRouter(r *mux.Router) {

	noteRouter := r.PathPrefix("/notes").Subrouter()
	noteRouter.HandleFunc("/", controller.NoteGetAllController).Methods(http.MethodGet)
	noteRouter.HandleFunc("/{tag}", controller.NoteGetByTagController).Methods(http.MethodGet)
	noteRouter.HandleFunc("/", controller.NoteCreateController).Methods(http.MethodPost)
}