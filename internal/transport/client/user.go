package client

import (
	"net"

	modelCl "github.com/DrIhor/test_project/internal/models/client"

	"github.com/DrIhor/test_project/internal/service/messages"
)

type user struct {
	conn net.Conn
	serv modelCl.UserServiceFunc
}

func NewUser(conn net.Conn, msg *messages.MsgService) *user {

	return &user{
		conn: conn,
		serv: messages.NewMsgService(),
	}
}
