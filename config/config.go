package config

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

type Files struct{}
type Movies struct{}
type Users struct {
	RedisClient RedisConfig
}
