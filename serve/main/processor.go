package main

import (
	"Massusers/common/message"
	"Massusers/serve/processes"
	"Massusers/serve/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//根据客户端发送的消息种类不同，决定调用哪个函数
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMessType:
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsesType:
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
func (this *Processor) process2() (err error) {
	for {
		//这读取数据包，封装成函数
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出来")
			} else {
				fmt.Println("readPkgerr=", err)
				return err
			}

		}
		fmt.Println("mes", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}

}
