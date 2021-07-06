package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
)

// message structure
type message struct {
	From string `json:"from"`
	Msg  string `json:"msg"`
}

// enter user name for chat identify
func (ms *message) getUserName() {
	fmt.Println("Enter username:")
	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	ms.From = name
}

// add new message from user to send
func (ms *message) AddMessage(text string) {
	ms.Msg = text
}

// output for user data
func (ms *message) PrintMessage() {
	fmt.Printf("User: %sMessage: %s\n", ms.From, ms.Msg)
}

// read information from server by other users
func readServer(conn net.Conn, wg *sync.WaitGroup) {
	var ms message

	// message capability
	recieveBuffer := make([]byte, 2048)
	for {
		// read info from connection
		read_len, err := conn.Read(recieveBuffer)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		// read struct
		request_right := recieveBuffer[:read_len]
		if err := json.Unmarshal(request_right, &ms); err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		// print user data
		ms.PrintMessage()
	}
}

// send information to other users
func writeServer(ms message, conn net.Conn, wg *sync.WaitGroup) {
	for {
		// read text from terminal to send
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		// add new msg
		ms.AddMessage(text)
		req, err := json.Marshal(ms)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		// send to other users
		conn.Write(req)
	}
}

// main logic of client
func StartWork() {
	// connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}

	var ms message
	ms.getUserName() // enter personal indentify name

	// sync gorutines to don`t close main
	// if wg is done - client work is end
	var wg sync.WaitGroup
	wg.Add(1)

	// read info from server and send data at same time
	go readServer(conn, &wg)
	go writeServer(ms, conn, &wg)

	wg.Wait()
	fmt.Println("End of work)))")
}
