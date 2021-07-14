package server

import (
	"encoding/json"
	"fmt"
	"net"

	conf "github.com/DrIhor/test_project/pkg/config"
	msg "github.com/DrIhor/test_project/pkg/message"
)

func sendAllData(obj msg.Message, usersConnections map[string]*userInfo) {
	// encode user data
	req, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
	}

	// send information for all users
	for _, user := range usersConnections {
		user.conn.Write(req)
	}
}

// work witl all connections
func handleConnection(conn net.Conn, usersConnections map[string]*userInfo) {
	var obj msg.Message
	receiveBuffer := make([]byte, 2048)

	for {
		read_len, err := conn.Read(receiveBuffer)
		if err != nil {
			fmt.Println(err)
			conn.Close()

			user := usersConnections[conn.RemoteAddr().String()]
			delete(usersConnections, conn.RemoteAddr().String())
			disconnectionMessage(&obj, user.userName, len(usersConnections))
			sendAllData(obj, usersConnections)

			break
		}

		request_right := receiveBuffer[:read_len]
		if err := json.Unmarshal(request_right, &obj); err != nil {
			fmt.Println(err)
			conn.Close()

			user := usersConnections[conn.RemoteAddr().String()]
			delete(usersConnections, conn.RemoteAddr().String())
			disconnectionMessage(&obj, user.userName, len(usersConnections))
			sendAllData(obj, usersConnections)

			break
		}

		if obj.UpdateName {
			us := usersConnections[conn.RemoteAddr().String()]

			updateNameMessage(&obj, us.userName, obj.From, len(usersConnections))
			us.updateName(obj.From)
		} else if obj.FirstConnection {
			us := usersConnections[conn.RemoteAddr().String()]

			connectionMessage(&obj, obj.From, len(usersConnections))
			us.updateName(obj.From)
		}

		sendAllData(obj, usersConnections)
	}
}

// start server work
func StartServer() {

	serverConf := conf.InitConfig()

	// init listen
	ln, err := net.Listen("tcp", serverConf.GetAggress())
	if err != nil {
		panic(err)
	}

	usersConnections := make(map[string]*userInfo) // all users
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		usersConnections[conn.RemoteAddr().String()] = addNewUser(conn) // save new user

		go handleConnection(conn, usersConnections) // handle all work with connection
	}
}
