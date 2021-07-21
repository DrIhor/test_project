package client

import (
	"strings"

	msg "github.com/DrIhor/test_project/internal/service/messages"
)

func userEvents(data *msg.MsgService, msg string) bool {
	wordsInRow := strings.Fields(msg)

	// if empty string we skip send data
	if len(wordsInRow) == 0 {
		return true
	}

	switch wordsInRow[0] {
	case "@changeName":
		if len(wordsInRow) >= 2 {
			data.UpdateUserName(wordsInRow[1])
		}
	default:
		data.AddMessage(msg)
	}

	return false
}
