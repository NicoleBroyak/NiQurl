package redishandler

import "github.com/go-redis/redis/v8"

func insertUser(userdata [5]string) {
	id, _ := RDB.Get(Ctx, "USER_COUNT").Float64()
	RDB.Incr(Ctx, "USER_COUNT")
	RDB.ZAdd(Ctx, "username")
	RDB.ZAdd(Ctx, "username", &redis.Z{Score: id, Member: userdata[0]})
	RDB.RPush(Ctx, "firstname", userdata[1])
	RDB.RPush(Ctx, "lastname", userdata[2])
	RDB.ZAdd(Ctx, "email", &redis.Z{Score: id, Member: userdata[3]})
	RDB.RPush(Ctx, "regdate", userdata[4])
}
