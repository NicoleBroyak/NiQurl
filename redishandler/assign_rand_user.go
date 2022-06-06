package redishandler

import (
	"math/rand"

	"time"

	"github.com/go-redis/redis/v8"
)

func AssignRandUser(RDB *redis.Client) string {
	rand.Seed(time.Now().UTC().UnixNano())
	uc, _ := RDB.Get(Ctx, "USER_COUNT").Int()
	randnum := int64(rand.Intn(uc))
	user := RDB.ZRange(Ctx, "username", randnum, randnum).Val()
	return user[0]
}
