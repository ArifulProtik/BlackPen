// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ArifulProtik/BlackPen/ent/love"
	"github.com/ArifulProtik/BlackPen/ent/user"
	"github.com/google/uuid"
)

// Love is the model entity for the Love schema.
type Love struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Noteid holds the value of the "noteid" field.
	Noteid uuid.UUID `json:"noteid,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LoveQuery when eager-loading is set.
	Edges      LoveEdges `json:"edges"`
	user_loves *uuid.UUID
}

// LoveEdges holds the relations/edges for other nodes in the graph.
type LoveEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LoveEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Love) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case love.FieldID:
			values[i] = new(sql.NullInt64)
		case love.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case love.FieldNoteid:
			values[i] = new(uuid.UUID)
		case love.ForeignKeys[0]: // user_loves
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Love", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Love fields.
func (l *Love) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case love.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case love.FieldNoteid:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field noteid", values[i])
			} else if value != nil {
				l.Noteid = *value
			}
		case love.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case love.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_loves", values[i])
			} else if value.Valid {
				l.user_loves = new(uuid.UUID)
				*l.user_loves = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Love entity.
func (l *Love) QueryUser() *UserQuery {
	return (&LoveClient{config: l.config}).QueryUser(l)
}

// Update returns a builder for updating this Love.
// Note that you need to call Love.Unwrap() before calling this method if this Love
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Love) Update() *LoveUpdateOne {
	return (&LoveClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Love entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Love) Unwrap() *Love {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Love is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Love) String() string {
	var builder strings.Builder
	builder.WriteString("Love(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", noteid=")
	builder.WriteString(fmt.Sprintf("%v", l.Noteid))
	builder.WriteString(", created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Loves is a parsable slice of Love.
type Loves []*Love

func (l Loves) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
