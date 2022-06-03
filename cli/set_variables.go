package cli

import (
	"fmt"
	"redis"
)

func SetVariables() {
	rdb := redis.RedisStart()
	defer rdb.Close()
	u, _ := rdb.Get(redis.Ctx, "SHORT_URL_LEN").Result()
	t, _ := rdb.Get(redis.Ctx, "USER_WAIT_TIME").Result()
	if u == "" || t == "" {
		rdb.Set(redis.Ctx, "USER_WAIT_TIME", 1000, 0)
		rdb.Set(redis.Ctx, "SHORT_URL_LEN", 4, 0)
		i, _ := rdb.Do(redis.Ctx, "ZCOUNT", "longurl", "-inf", "+inf").Result()
		j, _ := rdb.Do(redis.Ctx, "ZCOUNT", "username", "-inf", "+inf").Result()
		fmt.Println(i, j)
		rdb.Set(redis.Ctx, "url_count", i, 0)
		rdb.Set(redis.Ctx, "user_count", j, 0)
		Settings()
		fmt.Println("initializing database keys")
	}
}
