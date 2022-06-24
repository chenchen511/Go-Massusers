package process

import (
	"Massusers/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.Smses
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}
	info := fmt.Sprintf("用户id:%d,大家说的话：%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}
