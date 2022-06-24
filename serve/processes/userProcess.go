package processes

import (
	"Massusers/common/message"
	"Massusers/serve/model"
	"Massusers/serve/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}

}
func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}
	//发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		return
	}

}
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.RegisterMesType
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//处理登录
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginmes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginmes)
	if err != nil {
		fmt.Println("json.Unmarshal", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginMessType
	var loginResMes message.LoginResMes
	user, err := model.MyUserDao.Login(loginmes.UserId, loginmes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 500
			loginResMes.Error = "服务器内部出现错误"
		}
	} else {
		loginResMes.Code = 200
		this.UserId = loginmes.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser(loginmes.UserId)
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user)
	}
	// if loginmes.UserId == 100 && loginmes.UserPwd == "123456" {
	// 	loginResMes.Code = 200

	// } else {
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "用户不存在，请注册在使用"
	// }
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
