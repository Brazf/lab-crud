package config

type Config struct {
	ServerPort string `env:"SERVER_PORT" default:"8080"`

	
}
