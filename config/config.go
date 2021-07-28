package config

import "time"

// List of available database
const (
	DatabaseTypeSQL = "sql"
)

// Config .
type Config struct {
	App      *App      `json:"app"`
	Database *Database `json:"database"`
	HTTP     *HTTP     `json:"http"`
}

// App .
type App struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Timezone string `json:"timezone"`
	Env      string `json:"env"`
}

// Database .
type Database struct {
	Type string       `json:"type"`
	SQL  *SQLDatabase `json:"sql"`
}

// SQLDatabase .
type SQLDatabase struct {
	Driver          string `json:"driver"`
	Dsn             string `json:"dsn"`
	MaxIdleConn     int    `json:"max_idle_conn"`
	MaxOpenConn     int    `json:"max_open_conn"`
	ConnMaxLifetime int    `json:"conn_max_lifetime"`
}

// HTTP .
type HTTP struct {
	Enable           bool          `json:"enable"`
	Port             int           `json:"port"`
	GracefullTimeout time.Duration `json:"gracefull_timeout"`
}
