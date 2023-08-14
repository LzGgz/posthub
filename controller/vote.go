package controller

import (
	"posthub/dao/redis"
	"posthub/logic"
	"posthub/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func Vote(c *gin.Context) {
	userId := GetId(c)
	pv := new(model.ParamVote)
	if err := c.ShouldBindJSON(pv); err != nil {
		zap.L().Error("vote with invalid params", zap.Error(err))
		if err, ok := err.(validator.ValidationErrors); ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.Vote(userId, pv.PostID, pv.Direction); err != nil {
		if err == redis.ErrVoteTimeExpire {
			ResponseError(c, CodeVoteTimeExpire)
			return
		}
		ResponseError(c, CodeServerbusy)
		return
	}
	ResponseSuccess(c, nil)
}
