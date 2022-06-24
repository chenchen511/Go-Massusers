package model

import (
	"Massusers/common/message"
	"net"
)

//全局结构体
type CurUser struct {
	Conn net.Conn
	message.User
}
