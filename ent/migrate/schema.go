// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthsColumns holds the columns for the "auths" table.
	AuthsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "sessionid", Type: field.TypeUUID},
		{Name: "is_blocked", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// AuthsTable holds the schema information for the "auths" table.
	AuthsTable = &schema.Table{
		Name:       "auths",
		Columns:    AuthsColumns,
		PrimaryKey: []*schema.Column{AuthsColumns[0]},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "body", Type: field.TypeString, Size: 2147483647},
		{Name: "note_id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_comments", Type: field.TypeUUID, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NotesColumns holds the columns for the "notes" table.
	NotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "body", Type: field.TypeString, Size: 2147483647},
		{Name: "tags", Type: field.TypeJSON},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "f_image", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_notess", Type: field.TypeUUID, Nullable: true},
	}
	// NotesTable holds the schema information for the "notes" table.
	NotesTable = &schema.Table{
		Name:       "notes",
		Columns:    NotesColumns,
		PrimaryKey: []*schema.Column{NotesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notes_users_notess",
				Columns:    []*schema.Column{NotesColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "profile_pic", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthsTable,
		CommentsTable,
		NotesTable,
		UsersTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = UsersTable
	NotesTable.ForeignKeys[0].RefTable = UsersTable
}
