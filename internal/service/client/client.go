package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"

	msg "github.com/DrIhor/test_project/internal/service/messages"
	"github.com/DrIhor/test_project/internal/service/server/config"
)

// read information from server by other users
func readServer(conn net.Conn, wg *sync.WaitGroup) {

	// message capability
	recieveBuffer := make([]byte, 2048)
	for {
		var ms msg.MsgService

		// read info from connection
		read_len, err := conn.Read(recieveBuffer)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		// read struct
		request_right := recieveBuffer[:read_len]
		if err := json.Unmarshal(request_right, &ms.Data); err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		// print user data
		ms.PrintMessage()
	}
}

// send information to other users
func writeServer(ms msg.MsgService, conn net.Conn, wg *sync.WaitGroup) {
	for {
		// read text from terminal to send
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		emptyRow := userEvents(&ms, text) // create some event with data
		if emptyRow {
			continue
		}

		req, err := json.Marshal(ms.Data)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		// send to other users
		conn.Write(req)
	}
}

func FirstConnectionUpdate(ms *msg.MsgService, conn net.Conn) {
	ms.Data.FirstConnection = true // set updating of data

	req, err := json.Marshal(ms.Data)
	if err != nil {
		fmt.Println(err)
	}

	// send to other users
	conn.Write(req)

	ms.Data.FirstConnection = false // after end of update return to default value
}

// main logic of client
func StartWork() {

	server := config.InitServer()
	// connect to server
	conn, err := net.Dial("tcp", config.GetAggress(*server))
	if err != nil {
		panic(err)
	}

	var ms msg.MsgService
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
