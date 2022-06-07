package redishandler

import (
	"github.com/go-redis/redis/v8"
)

func CheckIntVar(val string, RDB *redis.Client) (int, error) {

	_, err := RDB.Do(Ctx, "ZCOUNT", val, "-inf", "+inf").Result()
	if err != nil {
		return 0, err
	}

	v, _ := RDB.Do(Ctx, "ZCOUNT", val, "-inf", "+inf").Int()
	RDB.Set(Ctx, val, v, 0)
	return v, nil
}
