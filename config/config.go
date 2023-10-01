package config

type CassandraConfig struct {
	Host     string
	Port     int
	Keyspace string
}

type PostgresConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type MySQLConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type SurrealDBConfig struct {
	Host      string
	Port      int
	Database  string
	Namespace string
	User      string
	Password  string
}

type Users struct {
	DatabaseConfig PostgresConfig
}
