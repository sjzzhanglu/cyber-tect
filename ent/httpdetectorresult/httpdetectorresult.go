// Code generated by entc, DO NOT EDIT.

package httpdetectorresult

import (
	"time"
)

const (
	// Label holds the string label denoting the httpdetectorresult type in the database.
	Label = "http_detector_result"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTask holds the string denoting the task field in the database.
	FieldTask = "task"
	// FieldResult holds the string denoting the result field in the database.
	FieldResult = "result"
	// FieldMaxDuration holds the string denoting the maxduration field in the database.
	FieldMaxDuration = "max_duration"
	// FieldMessages holds the string denoting the messages field in the database.
	FieldMessages = "messages"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldResults holds the string denoting the results field in the database.
	FieldResults = "results"

	// Table holds the table name of the httpdetectorresult in the database.
	Table = "http_detector_results"
)

// Columns holds all SQL columns for httpdetectorresult fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTask,
	FieldResult,
	FieldMaxDuration,
	FieldMessages,
	FieldURL,
	FieldResults,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// ResultValidator is a validator for the "result" field. It is called by the builders before save.
	ResultValidator func(int8) error
)
