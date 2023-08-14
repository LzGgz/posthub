package controller

import (
	"database/sql"
	"posthub/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityList(c *gin.Context) {
	data, err := logic.CommunityList()
	if err != nil {
		zap.L().Error("logic.CommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerbusy)
		return
	}
	ResponseSuccess(c, data)
}

func Community(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 32)
	if err != nil {
		zap.L().Error("get community with invalid id", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	comm, err := logic.CommunityById(id)
	if err != nil {
		zap.L().Error("logic.CommunityById failed", zap.Error(err))
		if err == sql.ErrNoRows {
			ResponseError(c, CodeNoSuchCommunityId)
			return
		}
		ResponseError(c, CodeServerbusy)
		return
	}
	ResponseSuccess(c, comm)
}
