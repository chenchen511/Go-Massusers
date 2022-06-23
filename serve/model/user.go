package model

type User struct {
	//为了保证序列化和反序列化成功，用户信息的json字符串的key和结构体的字段对应tag
	UserId   int    `json"UserId"`
	UserPwd  string `json"UserPwd"`
	UserName string `json"UserName"`
}
