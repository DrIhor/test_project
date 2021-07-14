package server

import (
	"fmt"

	msg "github.com/DrIhor/test_project/pkg/message"
)

func disconnectionMessage(user *msg.Message, userName string, usersInChat int) {
	user.AddMessage(fmt.Sprintf("\nAlert! %s is disconnected!!!\n  %d users in chat)", userName, usersInChat))
}

func connectionMessage(user *msg.Message, userName string, usersInChat int) {
	user.AddMessage(fmt.Sprintf("\nAlert! %s is connected!!! Say `Hello!!!\n %d users in chat)`", userName, usersInChat))
}

func updateNameMessage(user *msg.Message, oldName, newName string, usersInChat int) {
	user.AddMessage(fmt.Sprintf("\nAlert! %s change name to %s\n. %d users in chat)", oldName, newName, usersInChat))
}
