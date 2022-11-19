package config

import "os"

type Config struct {
	Port   string
	ConnDB string
}

func NewConfig() *Config {
	config := Config{}
	config.Port = os.Getenv("PORT")
	config.ConnDB = os.Getenv("CONN_DB")
	return &config

}
