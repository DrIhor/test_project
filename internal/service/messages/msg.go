package messages

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	msg "github.com/DrIhor/test_project/internal/models/message"
	"github.com/DrIhor/test_project/internal/service/server/users"
)

type MsgService struct {
	Data msg.Message
}

func (ms MsgService) DataEncode() ([]byte, error) {
	res, err := json.Marshal(ms.Data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// enter user name for chat identify
func (ms *MsgService) GetUserName() {
	fmt.Println("Enter username:")
	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	ms.Data.From = name
}

// add new message from user to send and add other default values
func (ms *MsgService) AddMessage(text string) {
	ms.Data.Msg = text // message text

	// other values to defaults
	ms.Data.UpdateName = false
	ms.Data.FirstConnection = false
}

func (ms *MsgService) UpdateUserName(name string) {
	fmt.Printf("\n Update user name to %s \n", name)
	ms.Data.UpdateName = true
	ms.Data.From = name
}

// output for user data
func (ms *MsgService) PrintMessage() {
	fmt.Printf("User: %s\nMessage: %s\n", ms.Data.From, ms.Data.Msg)
}

func (ms *MsgService) CheckUserUpdates(conn net.Conn, usServ *users.UserService) {
	if ms.Data.UpdateName {
		us := usServ.GetUser(conn.RemoteAddr().String())
		ms.UpdateNameMessage(us.UserName, ms.Data.From)
		usServ.Update(conn.RemoteAddr().String(), us.UserName)
	} else if ms.Data.FirstConnection {
		us := usServ.GetUser(conn.RemoteAddr().String())
		ms.ConnectionMessage(us.UserName, usServ.Count())
		usServ.Update(conn.RemoteAddr().String(), us.UserName)
	}
}
