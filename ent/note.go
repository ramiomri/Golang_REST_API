// Code generated by ent, DO NOT EDIT.

package ent

import (
	"example/golang-rest-api/ent/note"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Note is the model entity for the Note schema.
type Note struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Tag holds the value of the "tag" field.
	Tag string `json:"tag,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Note) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case note.FieldID:
			values[i] = new(sql.NullInt64)
		case note.FieldMessage, note.FieldTag:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Note", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Note fields.
func (n *Note) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case note.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case note.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				n.Message = value.String
			}
		case note.FieldTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag", values[i])
			} else if value.Valid {
				n.Tag = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Note.
// Note that you need to call Note.Unwrap() before calling this method if this Note
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Note) Update() *NoteUpdateOne {
	return (&NoteClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Note entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Note) Unwrap() *Note {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Note is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Note) String() string {
	var builder strings.Builder
	builder.WriteString("Note(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("message=")
	builder.WriteString(n.Message)
	builder.WriteString(", ")
	builder.WriteString("tag=")
	builder.WriteString(n.Tag)
	builder.WriteByte(')')
	return builder.String()
}

// Notes is a parsable slice of Note.
type Notes []*Note

func (n Notes) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
