package config

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

type Users struct {
	RedisClient RedisConfig
}
