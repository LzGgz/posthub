package controller

import (
	"posthub/logic"
	"posthub/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func CreatePost(c *gin.Context) {
	post := new(model.Post)
	if err := c.ShouldBindJSON(post); err != nil {
		zap.L().Error("create post with invalid params", zap.Error(err))
		if err, ok := err.(validator.ValidationErrors); ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	post.AuthorId = GetId(c)
	if err := logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerbusy)
		return
	}
	ResponseSuccess(c, nil)
}

func PostDetail(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		zap.L().Error("invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	pd, err := logic.PostDetailById(id)
	if err != nil {
		if err == logic.ErrInvalidPostId {
			ResponseError(c, CodeInvalidPostId)
			return
		}
		ResponseError(c, CodeServerbusy)
		return
	}
	ResponseSuccess(c, pd)
}

func Postlist(c *gin.Context) {
	pagestr := c.DefaultQuery("page", "0")
	pageNum, err := strconv.ParseInt(pagestr, 10, 64)
	if err != nil {
		zap.L().Error("invalid page", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	pds, err := logic.PostList(int(pageNum))
	if err != nil {
		ResponseError(c, CodeServerbusy)
		return
	}
	ResponseSuccess(c, pds)
}

func Posts(c *gin.Context) {
	p := new(model.ParamPostList)
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("get post list with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	pds, err := logic.Posts(p)
	if err != nil {
		ResponseError(c, CodeServerbusy)
		return
	}
	ResponseSuccess(c, pds)
}
