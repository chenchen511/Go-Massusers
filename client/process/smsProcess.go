package process

import (
	"Massusers/common/message"
	"Massusers/serve/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsesType
	var smsMes message.Smses
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes err", err.Error())
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes err", err.Error())
	}
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("senggroup", err)
		return
	}
	return
}
