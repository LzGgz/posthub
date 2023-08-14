package controller

import (
	"database/sql"
	"posthub/dao/mysql"
	"posthub/logic"
	"posthub/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	p := new(model.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid params", zap.Error(err))
		if err, ok := err.(validator.ValidationErrors); ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))
		if err == mysql.ErrorUsernameExist {
			ResponseError(c, CodeUsernameExist)
			return
		}
		ResponseError(c, CodeServerbusy)
		return
	}
	ResponseSuccess(c, nil)
}

func Login(c *gin.Context) {
	p := new(model.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid params", zap.Error(err))
		if err, ok := err.(validator.ValidationErrors); ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	jwt, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if err == sql.ErrNoRows {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerbusy)
		return
	}
	ResponseSuccess(c, jwt)
}
