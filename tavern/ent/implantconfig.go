// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/kcarretto/realm/tavern/ent/implantconfig"
)

// ImplantConfig is the model entity for the ImplantConfig schema.
type ImplantConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	// Human-readable name assigned to this implant config (e.g. imix-rhel).
	Name string `json:"name,omitempty"`
	// AuthToken holds the value of the "authToken" field.
	// Authentication token used by the implant to communicate with the GraphQL API.
	AuthToken string `json:"authToken,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImplantConfigQuery when eager-loading is set.
	Edges ImplantConfigEdges `json:"edges"`
}

// ImplantConfigEdges holds the relations/edges for other nodes in the graph.
type ImplantConfigEdges struct {
	// DeploymentConfigs holds the value of the deploymentConfigs edge.
	DeploymentConfigs []*DeploymentConfig `json:"deploymentConfigs,omitempty"`
	// Implants holds the value of the implants edge.
	Implants []*Implant `json:"implants,omitempty"`
	// ServiceConfigs holds the value of the serviceConfigs edge.
	ServiceConfigs []*ImplantServiceConfig `json:"serviceConfigs,omitempty"`
	// CallbackConfigs holds the value of the callbackConfigs edge.
	CallbackConfigs []*ImplantCallbackConfig `json:"callbackConfigs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// DeploymentConfigsOrErr returns the DeploymentConfigs value or an error if the edge
// was not loaded in eager-loading.
func (e ImplantConfigEdges) DeploymentConfigsOrErr() ([]*DeploymentConfig, error) {
	if e.loadedTypes[0] {
		return e.DeploymentConfigs, nil
	}
	return nil, &NotLoadedError{edge: "deploymentConfigs"}
}

// ImplantsOrErr returns the Implants value or an error if the edge
// was not loaded in eager-loading.
func (e ImplantConfigEdges) ImplantsOrErr() ([]*Implant, error) {
	if e.loadedTypes[1] {
		return e.Implants, nil
	}
	return nil, &NotLoadedError{edge: "implants"}
}

// ServiceConfigsOrErr returns the ServiceConfigs value or an error if the edge
// was not loaded in eager-loading.
func (e ImplantConfigEdges) ServiceConfigsOrErr() ([]*ImplantServiceConfig, error) {
	if e.loadedTypes[2] {
		return e.ServiceConfigs, nil
	}
	return nil, &NotLoadedError{edge: "serviceConfigs"}
}

// CallbackConfigsOrErr returns the CallbackConfigs value or an error if the edge
// was not loaded in eager-loading.
func (e ImplantConfigEdges) CallbackConfigsOrErr() ([]*ImplantCallbackConfig, error) {
	if e.loadedTypes[3] {
		return e.CallbackConfigs, nil
	}
	return nil, &NotLoadedError{edge: "callbackConfigs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ImplantConfig) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case implantconfig.FieldID:
			values[i] = new(sql.NullInt64)
		case implantconfig.FieldName, implantconfig.FieldAuthToken:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ImplantConfig", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ImplantConfig fields.
func (ic *ImplantConfig) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case implantconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ic.ID = int(value.Int64)
		case implantconfig.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ic.Name = value.String
			}
		case implantconfig.FieldAuthToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authToken", values[i])
			} else if value.Valid {
				ic.AuthToken = value.String
			}
		}
	}
	return nil
}

// QueryDeploymentConfigs queries the "deploymentConfigs" edge of the ImplantConfig entity.
func (ic *ImplantConfig) QueryDeploymentConfigs() *DeploymentConfigQuery {
	return (&ImplantConfigClient{config: ic.config}).QueryDeploymentConfigs(ic)
}

// QueryImplants queries the "implants" edge of the ImplantConfig entity.
func (ic *ImplantConfig) QueryImplants() *ImplantQuery {
	return (&ImplantConfigClient{config: ic.config}).QueryImplants(ic)
}

// QueryServiceConfigs queries the "serviceConfigs" edge of the ImplantConfig entity.
func (ic *ImplantConfig) QueryServiceConfigs() *ImplantServiceConfigQuery {
	return (&ImplantConfigClient{config: ic.config}).QueryServiceConfigs(ic)
}

// QueryCallbackConfigs queries the "callbackConfigs" edge of the ImplantConfig entity.
func (ic *ImplantConfig) QueryCallbackConfigs() *ImplantCallbackConfigQuery {
	return (&ImplantConfigClient{config: ic.config}).QueryCallbackConfigs(ic)
}

// Update returns a builder for updating this ImplantConfig.
// Note that you need to call ImplantConfig.Unwrap() before calling this method if this ImplantConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (ic *ImplantConfig) Update() *ImplantConfigUpdateOne {
	return (&ImplantConfigClient{config: ic.config}).UpdateOne(ic)
}

// Unwrap unwraps the ImplantConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ic *ImplantConfig) Unwrap() *ImplantConfig {
	tx, ok := ic.config.driver.(*txDriver)
	if !ok {
		panic("ent: ImplantConfig is not a transactional entity")
	}
	ic.config.driver = tx.drv
	return ic
}

// String implements the fmt.Stringer.
func (ic *ImplantConfig) String() string {
	var builder strings.Builder
	builder.WriteString("ImplantConfig(")
	builder.WriteString(fmt.Sprintf("id=%v", ic.ID))
	builder.WriteString(", name=")
	builder.WriteString(ic.Name)
	builder.WriteString(", authToken=")
	builder.WriteString(ic.AuthToken)
	builder.WriteByte(')')
	return builder.String()
}

// ImplantConfigs is a parsable slice of ImplantConfig.
type ImplantConfigs []*ImplantConfig

func (ic ImplantConfigs) config(cfg config) {
	for _i := range ic {
		ic[_i].config = cfg
	}
}