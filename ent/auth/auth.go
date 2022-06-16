// Code generated by entc, DO NOT EDIT.

package auth

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the auth type in the database.
	Label = "auth"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSessionID holds the string denoting the session_id field in the database.
	FieldSessionID = "session_id"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldIsBlocked holds the string denoting the is_blocked field in the database.
	FieldIsBlocked = "is_blocked"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the auth in the database.
	Table = "auths"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "auths"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_authentication"
)

// Columns holds all SQL columns for auth fields.
var Columns = []string{
	FieldID,
	FieldSessionID,
	FieldIP,
	FieldIsBlocked,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "auths"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_authentication",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// IPValidator is a validator for the "ip" field. It is called by the builders before save.
	IPValidator func(string) error
	// DefaultIsBlocked holds the default value on creation for the "is_blocked" field.
	DefaultIsBlocked bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
