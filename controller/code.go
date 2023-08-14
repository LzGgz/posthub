package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Code int

const (
	CodeSuccess Code = 1000 + iota
	CodeInvalidParam
	CodeUsernameExist
	CodeUserNotExist
	CodeServerbusy

	CodeNeedLogin
	CodeInvalidToken
	CodeTokenInvaildFormat

	CodeNoSuchCommunityId

	CodeInvalidPostId

	CodeVoteTimeExpire
)

var codeMsgs = map[Code]string{
	CodeSuccess:       "success",
	CodeInvalidParam:  "无效参数",
	CodeUsernameExist: "用户名已存在",
	CodeUserNotExist:  "用户名或密码错误",
	CodeServerbusy:    "服务器繁忙",

	CodeNeedLogin:          "需要登陆",
	CodeInvalidToken:       "无效的Token",
	CodeTokenInvaildFormat: "Token格式有误",

	CodeNoSuchCommunityId: "社区id无效",
	CodeInvalidPostId:     "帖子id无效",

	CodeVoteTimeExpire: "投票截止日期已过",
}

func (c Code) GetMsg() string {
	msg, ok := codeMsgs[c]
	if !ok {
		msg = codeMsgs[CodeServerbusy]
	}
	return msg
}

func ResponseError(c *gin.Context, code Code) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  code.GetMsg(),
	})
}
func ResponseErrorWithMsg(c *gin.Context, code Code, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  CodeSuccess.GetMsg(),
		"data": data,
	})
}
