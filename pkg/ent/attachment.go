// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Skijetler/alphinium/pkg/ent/attachment"
	"github.com/Skijetler/alphinium/pkg/ent/post"
	"github.com/Skijetler/alphinium/pkg/ent/user"
)

// Attachment is the model entity for the Attachment schema.
type Attachment struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Link holds the value of the "link" field.
	Link string `json:"link,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Size holds the value of the "size" field.
	Size string `json:"size,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint64 `json:"user_id,omitempty"`
	// PostID holds the value of the "post_id" field.
	PostID uint64 `json:"post_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttachmentQuery when eager-loading is set.
	Edges AttachmentEdges `json:"edges"`
}

// AttachmentEdges holds the relations/edges for other nodes in the graph.
type AttachmentEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Post holds the value of the post edge.
	Post *Post `json:"post,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttachmentEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttachmentEdges) PostOrErr() (*Post, error) {
	if e.loadedTypes[1] {
		if e.Post == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attachment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attachment.FieldID, attachment.FieldUserID, attachment.FieldPostID:
			values[i] = new(sql.NullInt64)
		case attachment.FieldLink, attachment.FieldName, attachment.FieldSize, attachment.FieldType, attachment.FieldHash:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Attachment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attachment fields.
func (a *Attachment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attachment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case attachment.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				a.Link = value.String
			}
		case attachment.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case attachment.FieldSize:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				a.Size = value.String
			}
		case attachment.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = value.String
			}
		case attachment.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				a.Hash = value.String
			}
		case attachment.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				a.UserID = uint64(value.Int64)
			}
		case attachment.FieldPostID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field post_id", values[i])
			} else if value.Valid {
				a.PostID = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Attachment entity.
func (a *Attachment) QueryUser() *UserQuery {
	return (&AttachmentClient{config: a.config}).QueryUser(a)
}

// QueryPost queries the "post" edge of the Attachment entity.
func (a *Attachment) QueryPost() *PostQuery {
	return (&AttachmentClient{config: a.config}).QueryPost(a)
}

// Update returns a builder for updating this Attachment.
// Note that you need to call Attachment.Unwrap() before calling this method if this Attachment
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attachment) Update() *AttachmentUpdateOne {
	return (&AttachmentClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Attachment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attachment) Unwrap() *Attachment {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attachment is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attachment) String() string {
	var builder strings.Builder
	builder.WriteString("Attachment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("link=")
	builder.WriteString(a.Link)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(a.Size)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(a.Type)
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(a.Hash)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteString(", ")
	builder.WriteString("post_id=")
	builder.WriteString(fmt.Sprintf("%v", a.PostID))
	builder.WriteByte(')')
	return builder.String()
}

// Attachments is a parsable slice of Attachment.
type Attachments []*Attachment

func (a Attachments) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}