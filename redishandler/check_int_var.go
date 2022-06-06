package redishandler

import (
	"github.com/go-redis/redis/v8"
)

func CheckIntVar(val string, RDB *redis.Client) (int, error) {

	v, err := RDB.Do(Ctx, "ZCOUNT", val, "-inf", "+inf").Int()
	if err != nil {
		return 0, err
	}

	RDB.Set(Ctx, val, v, 0)
	return v, nil
}
