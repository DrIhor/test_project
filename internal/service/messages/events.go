package messages

import "strings"

func (sr *MsgService) CheckEvent(msg string) bool {
	wordsInRow := strings.Fields(msg)

	// if empty string we skip send data
	if len(wordsInRow) == 0 {
		return true
	}

	switch wordsInRow[0] {
	case "@changeName":
		if len(wordsInRow) >= 2 {
			sr.UpdateUserName(wordsInRow[1])
		}
	default:
		sr.AddMessage(msg)
	}

	return false
}
