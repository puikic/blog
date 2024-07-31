package test

import (
	"blog/database"
	"blog/util"
	"fmt"
	"testing"
)

func init() {
	util.InitLog("log")
}
func TestUser(t *testing.T) {
	user := database.GetUserByName("dqq")
	if user == nil {
		t.Fail()
		return
	}
	if user.PassWd != "e10adc3949ba59abbe56e057f20f883e" {
		fmt.Println(user.PassWd)
		t.Fail()
		return
	}
	user = database.GetUserByName("none")
	if user != nil {
		t.Fail()
		return
	}
}

func TestCreatUser(t *testing.T) {
	name, pass := "测试用户bot", "123456"
	database.CreateUser(name, pass)
}

func TestDelUser(t *testing.T) {
	database.DeleteUser("测试用户bot")
}

//go test -v ./database/test -run=^TestUser$ -count=1
//go test -v ./database/test -run=^TestCreatUser$ -count=1
//go test -v ./database/test -run=^TestDelUser$ -count=1
