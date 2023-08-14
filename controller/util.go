package controller

import (
	"github.com/gin-gonic/gin"
)

const CtxId = "id"

func GetId(c *gin.Context) int64 {
	id, _ := c.Get(CtxId)
	return id.(int64)
}
