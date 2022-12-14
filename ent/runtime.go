// Code generated by ent, DO NOT EDIT.

package ent

import (
	"example/golang-rest-api/ent/note"
	"example/golang-rest-api/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	noteFields := schema.Note{}.Fields()
	_ = noteFields
	// noteDescMessage is the schema descriptor for message field.
	noteDescMessage := noteFields[0].Descriptor()
	// note.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	note.MessageValidator = noteDescMessage.Validators[0].(func(string) error)
}
