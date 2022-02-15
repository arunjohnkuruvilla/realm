// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/kcarretto/realm/ent/target"
)

// Target is the model entity for the Target schema.
type Target struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	// A human readable identifier for the target system.
	Name string `json:"name,omitempty"`
	// ForwardConnectIP holds the value of the "forwardConnectIP" field.
	// The IP Address that can be used to connect to the target using a protocol like SSH.
	ForwardConnectIP string `json:"forwardConnectIP,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TargetQuery when eager-loading is set.
	Edges TargetEdges `json:"edges"`
}

// TargetEdges holds the relations/edges for other nodes in the graph.
type TargetEdges struct {
	// Credentials holds the value of the credentials edge.
	Credentials []*Credential `json:"credentials,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CredentialsOrErr returns the Credentials value or an error if the edge
// was not loaded in eager-loading.
func (e TargetEdges) CredentialsOrErr() ([]*Credential, error) {
	if e.loadedTypes[0] {
		return e.Credentials, nil
	}
	return nil, &NotLoadedError{edge: "credentials"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Target) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case target.FieldID:
			values[i] = new(sql.NullInt64)
		case target.FieldName, target.FieldForwardConnectIP:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Target", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Target fields.
func (t *Target) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case target.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case target.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case target.FieldForwardConnectIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field forwardConnectIP", values[i])
			} else if value.Valid {
				t.ForwardConnectIP = value.String
			}
		}
	}
	return nil
}

// QueryCredentials queries the "credentials" edge of the Target entity.
func (t *Target) QueryCredentials() *CredentialQuery {
	return (&TargetClient{config: t.config}).QueryCredentials(t)
}

// Update returns a builder for updating this Target.
// Note that you need to call Target.Unwrap() before calling this method if this Target
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Target) Update() *TargetUpdateOne {
	return (&TargetClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Target entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Target) Unwrap() *Target {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Target is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Target) String() string {
	var builder strings.Builder
	builder.WriteString("Target(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", forwardConnectIP=")
	builder.WriteString(t.ForwardConnectIP)
	builder.WriteByte(')')
	return builder.String()
}

// Targets is a parsable slice of Target.
type Targets []*Target

func (t Targets) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}