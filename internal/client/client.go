package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"

	conf "github.com/DrIhor/test_project/pkg/config"
	msg "github.com/DrIhor/test_project/pkg/message"
)

// read information from server by other users
func readServer(conn net.Conn, wg *sync.WaitGroup) {
	var ms msg.Message

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
func writeServer(ms msg.Message, conn net.Conn, wg *sync.WaitGroup) {
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

func FirstConnectionUpdate(ms *msg.Message, conn net.Conn) {
	ms.FirstConnection = true // set updating of data

	req, err := json.Marshal(ms)
	if err != nil {
		fmt.Println(err)
	}

	// send to other users
	conn.Write(req)

	ms.FirstConnection = false // after end of update return to default value
}

// main logic of client
func StartWork() {

	serverInfo := conf.InitConfig()
	// connect to server
	conn, err := net.Dial("tcp", serverInfo.GetAggress())
	if err != nil {
		panic(err)
	}

	var ms msg.Message
	ms.GetUserName() // enter personal indentify name

	// send user name to save at serever
	FirstConnectionUpdate(&ms, conn)

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
