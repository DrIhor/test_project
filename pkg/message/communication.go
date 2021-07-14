package model

import (
	"bufio"
	"fmt"
	"os"
)

// message structure between client and server
type Message struct {
	From            string `json:"from"`
	Msg             string `json:"msg"`
	UpdateName      bool   `json:"updateName"`
	FirstConnection bool   `jsob:"firstConnection"`
}

// enter user name for chat identify
func (ms *Message) GetUserName() {
	fmt.Println("Enter username:")
	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	ms.From = name
}

// add new message from user to send and add other default values
func (ms *Message) AddMessage(text string) {
	ms.Msg = text // message text

	// other values to defaults
	ms.UpdateName = false
	ms.FirstConnection = false
}

func (ms *Message) UpdateUserName(name string) {
	fmt.Printf("\n Update user name to %s \n", name)
	ms.UpdateName = true
	ms.From = name
}

// output for user data
func (ms *Message) PrintMessage() {
	fmt.Printf("User: %s\nMessage: %s\n", ms.From, ms.Msg)
}
