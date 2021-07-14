package model

import (
	"bufio"
	"fmt"
	"os"
)

// message structure between client and server
type Message struct {
	From string `json:"from"`
	Msg  string `json:"msg"`
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

// add new message from user to send
func (ms *Message) AddMessage(text string) {
	ms.Msg = text
}

// output for user data
func (ms *Message) PrintMessage() {
	fmt.Printf("User: %sMessage: %s\n", ms.From, ms.Msg)
}
