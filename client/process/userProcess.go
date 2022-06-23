package process

import (
	"Massusers/common/message"
	"Massusers/serve/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dail err=", err)
		return
	}
	//延时关闭
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("err", err)
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn),err", err)
		return

	}
	var registerResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		//	fmt.Println("登录成功")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// fmt.Printf("userId =%d userPwd=%s\n", userId, userPwd)
	// return nil

	//1.
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dail err=", err)
		return
	}
	//延时关闭
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMessType
	var LoginMes message.LoginMes
	LoginMes.UserId = userId
	LoginMes.UserPwd = userPwd
	data, err := json.Marshal(LoginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//var pkglen uint32
	pkglen := uint32(len(data))
	var bytes []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(bytes[0:4], pkglen)
	n, err := conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Printf("客户端发送消息的长度=%d内容=%s", len(data), string(data))
	//发送消息本身
	_, err = conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠20s")
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn),err", err)
		return

	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//	fmt.Println("登录成功")
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户ID", v)
		}
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
