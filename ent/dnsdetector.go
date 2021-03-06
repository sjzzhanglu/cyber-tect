// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/dnsdetector"
	"github.com/vicanso/cybertect/ent/schema"
)

// DNSDetector is the model entity for the DNSDetector schema.
type DNSDetector struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty" sql:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty" sql:"updated_at"`
	// Status holds the value of the "status" field.
	Status schema.Status `json:"status,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Receivers holds the value of the "receivers" field.
	Receivers []string `json:"receivers,omitempty"`
	// Timeout holds the value of the "timeout" field.
	Timeout string `json:"timeout,omitempty"`
	// Host holds the value of the "host" field.
	Host string `json:"host,omitempty"`
	// Ips holds the value of the "ips" field.
	Ips []string `json:"ips,omitempty"`
	// Servers holds the value of the "servers" field.
	Servers []string `json:"servers,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DNSDetector) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dnsdetector.FieldReceivers, dnsdetector.FieldIps, dnsdetector.FieldServers:
			values[i] = &[]byte{}
		case dnsdetector.FieldID, dnsdetector.FieldStatus:
			values[i] = &sql.NullInt64{}
		case dnsdetector.FieldName, dnsdetector.FieldOwner, dnsdetector.FieldDescription, dnsdetector.FieldTimeout, dnsdetector.FieldHost:
			values[i] = &sql.NullString{}
		case dnsdetector.FieldCreatedAt, dnsdetector.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type DNSDetector", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DNSDetector fields.
func (dd *DNSDetector) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dnsdetector.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dd.ID = int(value.Int64)
		case dnsdetector.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dd.CreatedAt = value.Time
			}
		case dnsdetector.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dd.UpdatedAt = value.Time
			}
		case dnsdetector.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				dd.Status = schema.Status(value.Int64)
			}
		case dnsdetector.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				dd.Name = value.String
			}
		case dnsdetector.FieldOwner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner", values[i])
			} else if value.Valid {
				dd.Owner = value.String
			}
		case dnsdetector.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				dd.Description = value.String
			}
		case dnsdetector.FieldReceivers:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field receivers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dd.Receivers); err != nil {
					return fmt.Errorf("unmarshal field receivers: %v", err)
				}
			}
		case dnsdetector.FieldTimeout:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timeout", values[i])
			} else if value.Valid {
				dd.Timeout = value.String
			}
		case dnsdetector.FieldHost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host", values[i])
			} else if value.Valid {
				dd.Host = value.String
			}
		case dnsdetector.FieldIps:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ips", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dd.Ips); err != nil {
					return fmt.Errorf("unmarshal field ips: %v", err)
				}
			}
		case dnsdetector.FieldServers:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field servers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dd.Servers); err != nil {
					return fmt.Errorf("unmarshal field servers: %v", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this DNSDetector.
// Note that you need to call DNSDetector.Unwrap() before calling this method if this DNSDetector
// was returned from a transaction, and the transaction was committed or rolled back.
func (dd *DNSDetector) Update() *DNSDetectorUpdateOne {
	return (&DNSDetectorClient{config: dd.config}).UpdateOne(dd)
}

// Unwrap unwraps the DNSDetector entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dd *DNSDetector) Unwrap() *DNSDetector {
	tx, ok := dd.config.driver.(*txDriver)
	if !ok {
		panic("ent: DNSDetector is not a transactional entity")
	}
	dd.config.driver = tx.drv
	return dd
}

// String implements the fmt.Stringer.
func (dd *DNSDetector) String() string {
	var builder strings.Builder
	builder.WriteString("DNSDetector(")
	builder.WriteString(fmt.Sprintf("id=%v", dd.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(dd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(dd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", dd.Status))
	builder.WriteString(", name=")
	builder.WriteString(dd.Name)
	builder.WriteString(", owner=")
	builder.WriteString(dd.Owner)
	builder.WriteString(", description=")
	builder.WriteString(dd.Description)
	builder.WriteString(", receivers=")
	builder.WriteString(fmt.Sprintf("%v", dd.Receivers))
	builder.WriteString(", timeout=")
	builder.WriteString(dd.Timeout)
	builder.WriteString(", host=")
	builder.WriteString(dd.Host)
	builder.WriteString(", ips=")
	builder.WriteString(fmt.Sprintf("%v", dd.Ips))
	builder.WriteString(", servers=")
	builder.WriteString(fmt.Sprintf("%v", dd.Servers))
	builder.WriteByte(')')
	return builder.String()
}

// DNSDetectors is a parsable slice of DNSDetector.
type DNSDetectors []*DNSDetector

func (dd DNSDetectors) config(cfg config) {
	for _i := range dd {
		dd[_i].config = cfg
	}
}
