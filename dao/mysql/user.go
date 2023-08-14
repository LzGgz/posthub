package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"posthub/model"
)

var secret = "ggz"
var (
	ErrorUsernameExist = errors.New("该用户名已存在")
	ErrorUserNotExist  = errors.New("用户名或密码错误")
)

func IsUsernameExist(username string) (err error) {
	sqlstr := "select count(user_id) from user where username = ?"
	var count int
	if err = db.Get(&count, sqlstr, username); err != nil {
		return
	}
	if count != 0 {
		err = ErrorUsernameExist
		return
	}
	return

}
func InsertUser(user *model.User) (err error) {
	sqlstr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err = db.Exec(sqlstr, user.UserID, user.Username, encyptString(user.Password))
	return

}
func encyptString(str string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(str)))
}

func Login(username string, password string) (id int64, err error) {
	sqlstr := "select user_id from user where username = ? and password = ?"
	err = db.Get(&id, sqlstr, username, encyptString(password))
	return
}

func UsernameById(id int64) (username string, err error) {
	sqlstr := "select username from user where user_id = ?"
	err = db.Get(&username, sqlstr, id)
	return
}
