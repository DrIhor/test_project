package server

import (
	"encoding/json"
	"fmt"
	"net"

	srModel "github.com/DrIhor/test_project/internal/models/server"

	"github.com/DrIhor/test_project/internal/service/server/config"
	"github.com/DrIhor/test_project/internal/service/server/users"

	msg "github.com/DrIhor/test_project/internal/service/messages"
)

// send message to each user
func sendAllData(data []byte, users []srModel.UserInfo) {
	// send information for all users
	for _, user := range users {
		user.Conn.Write(data)
	}
}

// work witl all connections
func handleConnection(conn net.Conn, usServ *users.UserService) {
	receiveBuffer := make([]byte, 2048)

	for {
		var message msg.MsgService

		read_len, err := conn.Read(receiveBuffer)
		if err != nil {
			fmt.Println(err)
			conn.Close()

			// remove user and send alert
			user := usServ.GetUser(conn.RemoteAddr().String())
			usersLeft, _ := usServ.Delete(conn.RemoteAddr().String())
			message.DisconnectionMessage(user.UserName, usersLeft)

			resp, _ := message.DataEncode()
			sendAllData(resp, usServ.GetAllUsers())

			break
		}

		request_right := receiveBuffer[:read_len]

		if err := json.Unmarshal(request_right, &message.Data); err != nil {
			fmt.Println(err)
			conn.Close()

			// remove user and send alert
			user := usServ.GetUser(conn.RemoteAddr().String())
			usersLeft, _ := usServ.Delete(conn.RemoteAddr().String())
			message.DisconnectionMessage(user.UserName, usersLeft)

			resp, _ := message.DataEncode()
			sendAllData(resp, usServ.GetAllUsers())

			break
		}

		message.CheckUserUpdates(conn, usServ)

		dataResponse, _ := message.DataEncode()
		sendAllData(dataResponse, usServ.GetAllUsers())
	}
}

// start server work
func StartServer() {

	server := config.InitServer()
	// init listen
	ln, err := net.Listen("tcp", config.GetAggress(*server))
	if err != nil {
		panic(err)
	}

	userServices := users.NewUserService(server.Storage)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		userServices.Add(conn.RemoteAddr().String(), "", conn)
		// usersConnections[conn.RemoteAddr().String()] = addNewUser(conn) // save new user

		go handleConnection(conn, userServices) // handle all work with connection
	}
}