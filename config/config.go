package config

type PostgresConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type Users struct {
	DatabaseConfig PostgresConfig
}
