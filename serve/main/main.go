package main

import (
	"Massusers/serve/model"
	_ "errors"
	"fmt"
	"net"
	"time"
)

// func readPkg(conn net.Conn) (mes message.Message, err error) {
// 	buf := make([]byte, 8096)
// 	fmt.Println("读取客户端发送的数据...")
// 	n, err := conn.Read(buf[:4])
// 	if n != 4 || err != nil {
// 		//err = errors.New("conn.Read err")
// 		return
// 	}
// 	var pkglen uint32
// 	pkglen = binary.BigEndian.Uint32(buf[0:4])
// 	n, err = conn.Read(buf[:pkglen])
// 	if n != int(pkglen) || err != nil {
// 		fmt.Println("conn.Read err=", err)
// 		return
// 	}
// 	json.Unmarshal(buf[:pkglen], &mes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal", err)
// 		return
// 	}
// 	return
// }
// func writePkg(conn net.Conn, data []byte) (err error) {
// 	pkglen := uint32(len(data))
// 	var bytes []byte = make([]byte, 4)
// 	binary.BigEndian.PutUint32(bytes[0:4], pkglen)
// 	n, err := conn.Write(bytes)
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(bytes) fail", err)
// 		return
// 	}
// 	n, err = conn.Write(data)
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(bytes) fail", err)
// 		return
// 	}
// 	return
// }

// //处理登录
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	var loginmes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginmes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal", err)
// 		return
// 	}
// 	var resMes message.Message
// 	resMes.Type = message.LoginMessType
// 	var loginResMes message.LoginResMes
// 	if loginmes.UserId == 100 && loginmes.UserPwd == "123456" {
// 		loginResMes.Code = 200

// 	} else {
// 		loginResMes.Code = 500
// 		loginResMes.Error = "用户不存在，请注册在使用"
// 	}
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal", err)
// 		return
// 	}
// 	resMes.Data = string(data)
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal", err)
// 		return
// 	}
// 	err = writePkg(conn, data)
// 	return
// }

// //根据客户端发送的消息种类不同，决定调用哪个函数
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMessType:
// 		err = serverProcessLogin(conn, mes)
// 	case message.RegisterType:
// 	default:
// 		fmt.Println("消息类型不存在，无法处理")
// 	}
// 	return
// }

//处理客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()
	// buf := make([]byte, 8096)
	process := &Processor{
		Conn: conn,
	}
	err := process.process2()
	if err != nil {
		fmt.Println("客户端出错", err)
		return
	}
}

//完成对myUserDao的初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	//初始化连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	for {
		fmt.Println("等待客户端来链接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen Accept err=", err)
		}
		go process(conn)
	}
}
