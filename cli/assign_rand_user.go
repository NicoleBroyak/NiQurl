package cli

import (
	"fmt"
	"math/rand"
	"redis"
	"strconv"
	"time"
)

func AssignRandUser() string {
	rand.Seed(time.Now().UTC().UnixNano())
	rdb := redis.RedisStart()
	defer rdb.Close()
	uc, _ := rdb.Get(redis.Ctx, "user_count").Result()
	u, _ := strconv.Atoi(uc)
	randnum := rand.Intn(u)
	fmt.Println(randnum)
	username, _ := rdb.Do(redis.Ctx, "ZRANGE", "username", randnum, randnum).Result()
	fmt.Println(username)
	user := fmt.Sprintf("%v", username)
	user = string(user[1 : len(user)-1])
	return user
}
