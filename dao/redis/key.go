package redis

const (
	keyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"
	KeyPostScoreZSet   = "post:score"
	KeyPostVotedZSetPF = "post:voted:"
)

func RedisKey(key string) string {
	return keyPrefix + key
}
