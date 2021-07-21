package messages

import (
	"fmt"
)

func (ms *MsgService) DisconnectionMessage(userName string, usersInChat int) {
	ms.Data.Msg = fmt.Sprintf("\nAlert! %s is disconnected!!!\n  %d users in chat)", userName, usersInChat)
}

func (ms *MsgService) ConnectionMessage(userName string, usersInChat int) {
	ms.Data.Msg = fmt.Sprintf("\nAlert! %s is connected!!! Say `Hello!!!\n %d users in chat)`", userName, usersInChat)
}

func (ms *MsgService) UpdateNameMessage(oldName, newName string) {
	ms.Data.Msg = fmt.Sprintf("\nAlert! %s change name to %s\n.)", oldName, newName)
}
