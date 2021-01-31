// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// ConfigurationsColumns holds the columns for the "configurations" table.
	ConfigurationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "category", Type: field.TypeEnum, Enums: []string{"mockTime", "blockIP", "signedKey", "routerConcurrency", "router", "sessionInterceptor"}},
		{Name: "owner", Type: field.TypeString},
		{Name: "data", Type: field.TypeString},
		{Name: "started_at", Type: field.TypeTime},
		{Name: "ended_at", Type: field.TypeTime},
	}
	// ConfigurationsTable holds the schema information for the "configurations" table.
	ConfigurationsTable = &schema.Table{
		Name:        "configurations",
		Columns:     ConfigurationsColumns,
		PrimaryKey:  []*schema.Column{ConfigurationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "configuration_name",
				Unique:  true,
				Columns: []*schema.Column{ConfigurationsColumns[4]},
			},
			{
				Name:    "configuration_status",
				Unique:  false,
				Columns: []*schema.Column{ConfigurationsColumns[3]},
			},
		},
	}
	// DNSDetectorsColumns holds the columns for the "dns_detectors" table.
	DNSDetectorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "owner", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "receivers", Type: field.TypeJSON},
		{Name: "timeout", Type: field.TypeString},
		{Name: "host", Type: field.TypeString},
		{Name: "ips", Type: field.TypeJSON},
		{Name: "servers", Type: field.TypeJSON},
	}
	// DNSDetectorsTable holds the schema information for the "dns_detectors" table.
	DNSDetectorsTable = &schema.Table{
		Name:        "dns_detectors",
		Columns:     DNSDetectorsColumns,
		PrimaryKey:  []*schema.Column{DNSDetectorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "dnsdetector_owner",
				Unique:  false,
				Columns: []*schema.Column{DNSDetectorsColumns[5]},
			},
		},
	}
	// DNSDetectorResultsColumns holds the columns for the "dns_detector_results" table.
	DNSDetectorResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "task", Type: field.TypeInt},
		{Name: "result", Type: field.TypeInt8},
		{Name: "max_duration", Type: field.TypeInt},
		{Name: "messages", Type: field.TypeJSON},
		{Name: "host", Type: field.TypeString},
		{Name: "results", Type: field.TypeJSON},
	}
	// DNSDetectorResultsTable holds the schema information for the "dns_detector_results" table.
	DNSDetectorResultsTable = &schema.Table{
		Name:        "dns_detector_results",
		Columns:     DNSDetectorResultsColumns,
		PrimaryKey:  []*schema.Column{DNSDetectorResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "dnsdetectorresult_task",
				Unique:  false,
				Columns: []*schema.Column{DNSDetectorResultsColumns[4]},
			},
			{
				Name:    "dnsdetectorresult_result",
				Unique:  false,
				Columns: []*schema.Column{DNSDetectorResultsColumns[5]},
			},
		},
	}
	// HTTPDetectorsColumns holds the columns for the "http_detectors" table.
	HTTPDetectorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "owner", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "receivers", Type: field.TypeJSON},
		{Name: "timeout", Type: field.TypeString},
		{Name: "ips", Type: field.TypeJSON},
		{Name: "url", Type: field.TypeString},
	}
	// HTTPDetectorsTable holds the schema information for the "http_detectors" table.
	HTTPDetectorsTable = &schema.Table{
		Name:        "http_detectors",
		Columns:     HTTPDetectorsColumns,
		PrimaryKey:  []*schema.Column{HTTPDetectorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "httpdetector_owner",
				Unique:  false,
				Columns: []*schema.Column{HTTPDetectorsColumns[5]},
			},
		},
	}
	// HTTPDetectorResultsColumns holds the columns for the "http_detector_results" table.
	HTTPDetectorResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "task", Type: field.TypeInt},
		{Name: "result", Type: field.TypeInt8},
		{Name: "max_duration", Type: field.TypeInt},
		{Name: "messages", Type: field.TypeJSON},
		{Name: "url", Type: field.TypeString},
		{Name: "results", Type: field.TypeJSON},
	}
	// HTTPDetectorResultsTable holds the schema information for the "http_detector_results" table.
	HTTPDetectorResultsTable = &schema.Table{
		Name:        "http_detector_results",
		Columns:     HTTPDetectorResultsColumns,
		PrimaryKey:  []*schema.Column{HTTPDetectorResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "httpdetectorresult_task",
				Unique:  false,
				Columns: []*schema.Column{HTTPDetectorResultsColumns[4]},
			},
			{
				Name:    "httpdetectorresult_result",
				Unique:  false,
				Columns: []*schema.Column{HTTPDetectorResultsColumns[5]},
			},
		},
	}
	// PingDetectorsColumns holds the columns for the "ping_detectors" table.
	PingDetectorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "owner", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "receivers", Type: field.TypeJSON},
		{Name: "timeout", Type: field.TypeString},
		{Name: "ips", Type: field.TypeJSON},
	}
	// PingDetectorsTable holds the schema information for the "ping_detectors" table.
	PingDetectorsTable = &schema.Table{
		Name:        "ping_detectors",
		Columns:     PingDetectorsColumns,
		PrimaryKey:  []*schema.Column{PingDetectorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "pingdetector_owner",
				Unique:  false,
				Columns: []*schema.Column{PingDetectorsColumns[5]},
			},
		},
	}
	// PingDetectorResultsColumns holds the columns for the "ping_detector_results" table.
	PingDetectorResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "task", Type: field.TypeInt},
		{Name: "result", Type: field.TypeInt8},
		{Name: "max_duration", Type: field.TypeInt},
		{Name: "messages", Type: field.TypeJSON},
		{Name: "ips", Type: field.TypeString},
		{Name: "results", Type: field.TypeJSON},
	}
	// PingDetectorResultsTable holds the schema information for the "ping_detector_results" table.
	PingDetectorResultsTable = &schema.Table{
		Name:        "ping_detector_results",
		Columns:     PingDetectorResultsColumns,
		PrimaryKey:  []*schema.Column{PingDetectorResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "pingdetectorresult_task",
				Unique:  false,
				Columns: []*schema.Column{PingDetectorResultsColumns[4]},
			},
			{
				Name:    "pingdetectorresult_result",
				Unique:  false,
				Columns: []*schema.Column{PingDetectorResultsColumns[5]},
			},
		},
	}
	// TCPDetectorsColumns holds the columns for the "tcp_detectors" table.
	TCPDetectorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "owner", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "receivers", Type: field.TypeJSON},
		{Name: "timeout", Type: field.TypeString},
		{Name: "addrs", Type: field.TypeJSON},
	}
	// TCPDetectorsTable holds the schema information for the "tcp_detectors" table.
	TCPDetectorsTable = &schema.Table{
		Name:        "tcp_detectors",
		Columns:     TCPDetectorsColumns,
		PrimaryKey:  []*schema.Column{TCPDetectorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "tcpdetector_owner",
				Unique:  false,
				Columns: []*schema.Column{TCPDetectorsColumns[5]},
			},
		},
	}
	// TCPDetectorResultsColumns holds the columns for the "tcp_detector_results" table.
	TCPDetectorResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "task", Type: field.TypeInt},
		{Name: "result", Type: field.TypeInt8},
		{Name: "max_duration", Type: field.TypeInt},
		{Name: "messages", Type: field.TypeJSON},
		{Name: "addrs", Type: field.TypeString},
		{Name: "results", Type: field.TypeJSON},
	}
	// TCPDetectorResultsTable holds the schema information for the "tcp_detector_results" table.
	TCPDetectorResultsTable = &schema.Table{
		Name:        "tcp_detector_results",
		Columns:     TCPDetectorResultsColumns,
		PrimaryKey:  []*schema.Column{TCPDetectorResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "tcpdetectorresult_task",
				Unique:  false,
				Columns: []*schema.Column{TCPDetectorResultsColumns[4]},
			},
			{
				Name:    "tcpdetectorresult_result",
				Unique:  false,
				Columns: []*schema.Column{TCPDetectorResultsColumns[5]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "account", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "roles", Type: field.TypeJSON, Nullable: true},
		{Name: "groups", Type: field.TypeJSON, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "user_account",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[4]},
			},
		},
	}
	// UserLoginsColumns holds the columns for the "user_logins" table.
	UserLoginsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "account", Type: field.TypeString},
		{Name: "user_agent", Type: field.TypeString, Nullable: true},
		{Name: "ip", Type: field.TypeString, Nullable: true},
		{Name: "track_id", Type: field.TypeString, Nullable: true},
		{Name: "session_id", Type: field.TypeString, Nullable: true},
		{Name: "x_forwarded_for", Type: field.TypeString, Nullable: true},
		{Name: "country", Type: field.TypeString, Nullable: true},
		{Name: "province", Type: field.TypeString, Nullable: true},
		{Name: "city", Type: field.TypeString, Nullable: true},
		{Name: "isp", Type: field.TypeString, Nullable: true},
	}
	// UserLoginsTable holds the schema information for the "user_logins" table.
	UserLoginsTable = &schema.Table{
		Name:        "user_logins",
		Columns:     UserLoginsColumns,
		PrimaryKey:  []*schema.Column{UserLoginsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "userlogin_account",
				Unique:  false,
				Columns: []*schema.Column{UserLoginsColumns[3]},
			},
			{
				Name:    "userlogin_created_at",
				Unique:  false,
				Columns: []*schema.Column{UserLoginsColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ConfigurationsTable,
		DNSDetectorsTable,
		DNSDetectorResultsTable,
		HTTPDetectorsTable,
		HTTPDetectorResultsTable,
		PingDetectorsTable,
		PingDetectorResultsTable,
		TCPDetectorsTable,
		TCPDetectorResultsTable,
		UsersTable,
		UserLoginsTable,
	}
)

func init() {
}
