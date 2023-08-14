package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func Vote(userId, postId string, dire float64) (err error) {
	//获取帖子发布时间
	postTime := rdb.ZScore(RedisKey(KeyPostTimeZSet), postId).Val()
	//判断投票时间是否过期
	if time.Now().Unix()-int64(postTime) > oneWeekInSeconds {
		err = ErrVoteTimeExpire
		return
	}
	//获取历史投票
	ov := rdb.ZScore(RedisKey(KeyPostVotedZSetPF+postId), userId).Val()
	p := rdb.Pipeline()
	//重新计算帖子分数
	p.ZIncrBy(RedisKey(KeyPostScoreZSet), (dire-ov)*scorePerVote, postId)
	if dire == 0 {
		p.ZRem(RedisKey(KeyPostVotedZSetPF+postId), userId)
	} else {
		p.ZAdd(RedisKey(KeyPostVotedZSetPF+postId), redis.Z{
			Score:  dire,
			Member: userId,
		})
	}
	_, err = p.Exec()
	return
}

func CreatePost(postID int64) (err error) {
	p := rdb.TxPipeline()
	p.ZAdd(RedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	p.ZAdd(RedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err = p.Exec()
	return
}
func PostVotes(ids []string) (votes []int64, err error) {
	p := rdb.TxPipeline()
	for _, id := range ids {
		key := RedisKey(KeyPostVotedZSetPF + id)
		p.ZCount(key, "1", "1")
	}
	cmders, err := p.Exec()
	if err != nil {
		return
	}
	votes = make([]int64, 0, len(ids))
	for _, c := range cmders {
		v := c.(*redis.IntCmd).Val()
		votes = append(votes, v)
	}
	return
}
