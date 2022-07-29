package controller

import (
	"encoding/json"
	"example/golang-rest-api/ent"
	"example/golang-rest-api/service"
	"example/golang-rest-api/utils"
	"net/http"
	"github.com/gorilla/mux"
)


func NoteCreateController(w http.ResponseWriter, r *http.Request) {

	var newNote ent.Note
	err := json.NewDecoder(r.Body).Decode(&newNote)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	createdNote, err := service.NewNoteOps(r.Context()).NoteCreate(newNote)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
	}

	utils.Return(w, true, http.StatusOK, nil, createdNote)
}

func NoteGetAllController(w http.ResponseWriter, r *http.Request) {

	notes, err := service.NewNoteOps(r.Context()).NotesGetAll()
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, notes)
}

func NoteGetByTagController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tag:= vars["tag"]
	note, err := service.NewNoteOps(r.Context()).NoteGetByTag(tag)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, note)
}

