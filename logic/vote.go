package logic

import (
	"fmt"
	"posthub/dao/redis"

	"go.uber.org/zap"
)

func Vote(userId int64, postId string, dire int8) error {
	zap.L().Debug("VoteForPost", zap.Int64("userID", userId), zap.String("postID", postId), zap.Int8("dire", dire))
	return redis.Vote(fmt.Sprintf("%d", userId), postId, float64(dire))
}
