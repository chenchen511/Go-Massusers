package process

import (
	"Massusers/common/message"
	"Massusers/serve/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("-------恭喜登录成功")
	fmt.Println("-------1.显示在线用户列表")
	fmt.Println("-------2.发送消息")
	fmt.Println("-------3.信息列表")
	fmt.Println("-------4.退出系统")
	fmt.Println("-------请选择（1-4）")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("请输入要发送的话")
		fmt.Scanf("%d\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统")
		os.Exit(0)
	default:
		fmt.Println("你输入选项不正确")
	}
}
func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {
		fmt.Printf("客户端读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.Reading err=", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器端返回了未知消息类型")
		}
		fmt.Println(mes)
	}

}
