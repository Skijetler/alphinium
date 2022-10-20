// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Skijetler/alphinium/pkg/ent/category"
	"github.com/Skijetler/alphinium/pkg/ent/subcategory"
)

// Subcategory is the model entity for the Subcategory schema.
type Subcategory struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID uint64 `json:"category_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubcategoryQuery when eager-loading is set.
	Edges SubcategoryEdges `json:"edges"`
}

// SubcategoryEdges holds the relations/edges for other nodes in the graph.
type SubcategoryEdges struct {
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// Threads holds the value of the threads edge.
	Threads []*Thread `json:"threads,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubcategoryEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[0] {
		if e.Category == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// ThreadsOrErr returns the Threads value or an error if the edge
// was not loaded in eager-loading.
func (e SubcategoryEdges) ThreadsOrErr() ([]*Thread, error) {
	if e.loadedTypes[1] {
		return e.Threads, nil
	}
	return nil, &NotLoadedError{edge: "threads"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subcategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subcategory.FieldID, subcategory.FieldCategoryID:
			values[i] = new(sql.NullInt64)
		case subcategory.FieldName, subcategory.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subcategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subcategory fields.
func (s *Subcategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case subcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subcategory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case subcategory.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				s.CategoryID = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryCategory queries the "category" edge of the Subcategory entity.
func (s *Subcategory) QueryCategory() *CategoryQuery {
	return (&SubcategoryClient{config: s.config}).QueryCategory(s)
}

// QueryThreads queries the "threads" edge of the Subcategory entity.
func (s *Subcategory) QueryThreads() *ThreadQuery {
	return (&SubcategoryClient{config: s.config}).QueryThreads(s)
}

// Update returns a builder for updating this Subcategory.
// Note that you need to call Subcategory.Unwrap() before calling this method if this Subcategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subcategory) Update() *SubcategoryUpdateOne {
	return (&SubcategoryClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Subcategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subcategory) Unwrap() *Subcategory {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subcategory is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subcategory) String() string {
	var builder strings.Builder
	builder.WriteString("Subcategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", s.CategoryID))
	builder.WriteByte(')')
	return builder.String()
}

// Subcategories is a parsable slice of Subcategory.
type Subcategories []*Subcategory

func (s Subcategories) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}