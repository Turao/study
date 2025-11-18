package config

// CassandraConfig is the configuration for the Cassandra database
type CassandraConfig struct {
	Host     string
	Port     int
	Keyspace string
}

// PostgresConfig is the configuration for the Postgres database
type PostgresConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

// MySQLConfig is the configuration for the MySQL database
type MySQLConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

// SurrealDBConfig is the configuration for the SurrealDB database
type SurrealDBConfig struct {
	Host      string
	Port      int
	Database  string
	Namespace string
	User      string
	Password  string
}

// Users is the configuration for the Users service
type Users struct {
	DatabaseConfig PostgresConfig
}
