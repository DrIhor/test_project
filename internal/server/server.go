package server

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/DrIhor/test_project/internal/model"
)

// work witl all connections
func handleConnection(conn net.Conn, usersConnections map[string]net.Conn) {

	receiveBuffer := make([]byte, 2048)
	for {
		read_len, err := conn.Read(receiveBuffer)
		if err != nil {
			fmt.Println(err)
			conn.Close()
			delete(usersConnections, conn.RemoteAddr().String())
			fmt.Println(usersConnections)
			break
		}

		request_right := receiveBuffer[:read_len]

		var obj model.Message
		if err := json.Unmarshal(request_right, &obj); err != nil {
			fmt.Println(err)
			conn.Close()
			delete(usersConnections, conn.RemoteAddr().String())
			fmt.Println(usersConnections)
			break
		}

		// encode user data
		req, err := json.Marshal(obj)
		if err != nil {
			fmt.Println(err)
		}

		// send information for all users
		for _, cn := range usersConnections {
			cn.Write(req)
		}
	}

}

// start server work
func StartServer() {
	// init listen
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		panic(err)
	}

	usersConnections := make(map[string]net.Conn) // all users
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		usersConnections[conn.RemoteAddr().String()] = conn // save new user

		go handleConnection(conn, usersConnections) // handle all work with connection
	}
}
