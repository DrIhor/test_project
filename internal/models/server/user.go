package server

import "net"

type UserInfo struct {
	UserName string
	Conn     net.Conn
}

type UserService interface {
	Add(string, string, net.Conn) (int, error)
	Update(string, string) (bool, error)
	Count() int
	Delete(string) (int, error)
	GetUser(string) UserInfo
	GetAllUsers() []UserInfo
}
