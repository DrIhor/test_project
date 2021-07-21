package config

import (
	"fmt"
	"os"

	"github.com/DrIhor/test_project/internal/models/server"
	inMemory "github.com/DrIhor/test_project/internal/service/storage/imMemory"
)

type Server struct {
	Config  *server.ServerConfig
	Storage server.UserService
}

func InitServer() *Server {
	var server Server
	server.getServerConfig()
	server.getServerStorage()

	return &server
}

func (sr *Server) getServerConfig() {
	config := server.ServerConfig{
		Host: os.Getenv("SERVER_HOST"),
		Port: os.Getenv("SERVER_PORT"),
	}

	sr.Config = &config
}

func (sr *Server) getServerStorage() {
	stor := inMemory.UserData{}
	sr.Storage = &stor
}

func GetAggress(sr Server) string {
	return fmt.Sprintf("%s:%s", sr.Config.Host, sr.Config.Port)
}
