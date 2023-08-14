package redis

import (
	"posthub/model"
	"posthub/util/page"

	"github.com/spf13/viper"
)

const (
	OrderTime  = "time"
	OrderScore = "score"
)

func PostIDInOrder(p *model.ParamPostList) ([]string, error) {
	var key string
	if p.Order == OrderTime {
		key = RedisKey(KeyPostTimeZSet)
	}
	if p.Order == OrderScore {
		key = RedisKey(KeyPostScoreZSet)
	}
	start := page.Offset(p.Page)
	end := start + viper.GetInt("app.page_size") - 1
	return rdb.ZRevRange(key, int64(start), int64(end)).Result()
}
