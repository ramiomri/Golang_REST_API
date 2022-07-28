package service

import (
	"context"
	"example/golang-rest-api/config"
	"example/golang-rest-api/ent"
	"example/golang-rest-api/ent/note"
)

type NoteOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewNoteOps(ctx context.Context) *NoteOps {
	return &NoteOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *NoteOps) NotesGetAll() ([]*ent.Note, error) {

	notes, err := r.client.Note.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *NoteOps) NoteGetByTag(tag string) (*ent.Note, error) {

	notes, err := r.client.Note.Query().Where(note.Tag(tag)).Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *NoteOps) NoteCreate(newNote ent.Note) (*ent.Note, error) {

	newCreatedNote, err := r.client.Note.Create().
		SetMessage(newNote.Message).
		SetTag(newNote.Tag).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedNote, nil
}
