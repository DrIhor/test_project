package users

import (
	"net"

	"github.com/DrIhor/test_project/internal/models/server"
)

type UserService struct {
	storage server.UserService
}

func NewUserService(stor server.UserService) *UserService {
	return &UserService{
		storage: stor,
	}
}

func (sr *UserService) Add(address, userName string, conn net.Conn) (int, error) {
	return sr.storage.Add(address, userName, conn)
}

func (sr *UserService) Update(address, userName string) (bool, error) {
	return sr.storage.Update(address, userName)
}

func (sr *UserService) Count() int {
	return sr.storage.Count()
}

func (sr *UserService) Delete(adress string) (int, error) {
	return sr.storage.Delete(adress)
}

func (sr *UserService) GetUser(adress string) server.UserInfo {
	return sr.storage.GetUser(adress)
}

func (sr *UserService) GetAllUsers() []server.UserInfo {
	return sr.storage.GetAllUsers()
}
