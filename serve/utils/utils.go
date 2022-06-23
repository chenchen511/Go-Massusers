package utils

import (
	"Massusers/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//结构体
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		//err = errors.New("conn.Read err")
		return
	}
	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(this.Buf[0:4])
	n, err = this.Conn.Read(this.Buf[:pkglen])
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	json.Unmarshal(this.Buf[:pkglen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal", err)
		return
	}
	return
}
func (this *Transfer) WritePkg(data []byte) (err error) {
	pkglen := uint32(len(data))
	// var bytes []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(this.Buf[0:4], pkglen)
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	n, err = this.Conn.Write(data)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
