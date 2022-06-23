package model

import (
	"Massusers/common/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//初始化一个userDao实例
//
var (
	MyUserDao *UserDao
)

//使用工厂模式，创建一个userDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

type UserDao struct {
	pool *redis.Pool
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal,err=", err)
		return
	}
	return
}

//完成登陆校验
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		err = ERROR_USER_NOTEXISTS
		return
	}
	//这时证明用户获取到
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err != nil {
		err = ERROR_USER_EXISTS
		return
	}
	//这时说明id在redis还没有
	data, err := json.Marshal(user) //序列化
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存错误")
		return
	}
	return
}
