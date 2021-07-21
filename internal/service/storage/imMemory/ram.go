package inMemory

import (
	"net"

	userModel "github.com/DrIhor/test_project/internal/models/server"
)

type UserData userModel.UserInfo

var userList = map[string]*UserData{}

func InitStorage() *UserData {
	return &UserData{}
}

func (user *UserData) Add(address string, userName string, conn net.Conn) (int, error) {
	if _, ok := userList[address]; !ok {
		userList[address] = &UserData{
			UserName: userName,
			Conn:     conn,
		}
	}

	return len(userList), nil
}

func (user *UserData) Update(address, userName string) (bool, error) {
	if _, ok := userList[address]; ok {
		user := userList[address]
		user.UserName = userName
		return true, nil
	}

	return false, nil
}

func (user *UserData) Count() int {
	return len(userList)
}

func (user *UserData) Delete(address string) (int, error) {
	if _, ok := userList[address]; ok {
		delete(userList, address)
		return len(userList), nil
	}

	return len(userList), nil
}

func (user *UserData) GetUser(address string) userModel.UserInfo {
	obj := userModel.UserInfo(*userList[address])
	return obj
}

func (user *UserData) GetAllUsers() []userModel.UserInfo {
	var result []userModel.UserInfo
	for _, val := range userList {
		result = append(result, userModel.UserInfo(*val))
	}
	return result
}
