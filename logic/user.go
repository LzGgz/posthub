package logic

import (
	"posthub/dao/mysql"
	"posthub/model"
	"posthub/pkg/jwt"
	"posthub/pkg/snowflake"
)

func SignUp(p *model.ParamSignUp) (err error) {
	if err = mysql.IsUsernameExist(p.Username); err != nil {
		return
	}
	user := &model.User{
		UserID:   snowflake.GenID(),
		Username: p.Username,
		Password: p.Password,
	}
	err = mysql.InsertUser(user)
	return
}

func Login(p *model.ParamLogin) (string, error) {
	id, err := mysql.Login(p.Username, p.Password)
	if err != nil {
		return "", err
	}
	return jwt.GenToken(id)
}
