package server

import "net"

type userInfo struct {
	userName string
	conn     net.Conn
}

func addNewUser(conn net.Conn) *userInfo {
	return &userInfo{
		conn: conn,
	}
}

func (us *userInfo) updateName(name string) {
	us.userName = name
}
