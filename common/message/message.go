package message

const (
	LoginMessType           = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsesType               = "Smses"
)
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
type LoginResMes struct {
	Code    int `json:"code"`
	UsersId []int
	Error   string `json:"error"`
}
type RegisterMes struct {
	User User `json:"user"`
}
type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
type NotifyUserStatusMes struct {
	UserId int `json"userId"`
	Status int `json"status"`
}
type Smses struct {
	Content string `json:"content"`
	User           //匿名结构体，继承
}
