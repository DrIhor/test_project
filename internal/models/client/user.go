package client

import (
	"github.com/DrIhor/test_project/internal/service/messages"
)

type UserServiceFunc interface {
	CheckEvent(string) bool
	GetUserName()
	AddMessage(string)
	UpdateUserName(string)
	PrintMessage()
	GetMsgService() *messages.MsgService
}
