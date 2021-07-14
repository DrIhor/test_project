package config

import (
	"fmt"
	"os"
)

type serverConfig struct {
	host string
	port string
}

func InitConfig() *serverConfig {
	return &serverConfig{
		host: os.Getenv("SERVER_HOST"),
		port: os.Getenv("SERVER_PORT"),
	}
}

func (sr *serverConfig) GetAggress() string {
	return fmt.Sprintf("%s:%s", sr.host, sr.port)
}
