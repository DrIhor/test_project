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
func readServer(us *user, wg *sync.WaitGroup) {

	// message capability
	recieveBuffer := make([]byte, 2048)
	for {

		// read info from connection
		read_len, err := us.conn.Read(recieveBuffer)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		// read struct
		request_right := recieveBuffer[:read_len]
		msg := us.serv.GetMsgService()
		if err := json.Unmarshal(request_right, &msg.Data); err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		// print user data
		msg.PrintMessage()
	}
}

// send information to other users
func writeServer(us *user, wg *sync.WaitGroup) {
	for {
		// read text from terminal to send
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		emptyRow := us.serv.CheckEvent(text) // create some event with data
		if emptyRow {
			continue
		}

		msg := us.serv.GetMsgService()
		req, err := json.Marshal(msg.Data)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}

		// send to other users
		us.conn.Write(req)
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

	user := NewUser(conn, msg.NewMsgService())
	user.serv.GetUserName() // enter personal indentify name

	// send user name to save at serever
	msg := user.serv.GetMsgService()
	FirstConnectionUpdate(msg, conn)

	// sync gorutines to don`t close main
	// if wg is done - client work is end
	var wg sync.WaitGroup
	wg.Add(1)

	// read info from server and send data at same time
	go readServer(user, &wg)
	go writeServer(user, &wg)

	wg.Wait()
	fmt.Println("End of work)))")
}
