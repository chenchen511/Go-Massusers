package process

import (
	"Massusers/client/model"
	"Massusers/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

//显示当前在线用户
func outputOnlineUser() {
	fmt.Println("当前用户列表  ")
	for id, _ := range onlineUsers {
		fmt.Println("用户id", id)

	}
}

//处理返回信息
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId:     notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status

	onlineUsers[notifyUserStatusMes.UserId] = user
}
