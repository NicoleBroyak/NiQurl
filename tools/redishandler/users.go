package redishandler

import (
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nicolebroyak/niqurl/tools/randomusers"
)

func RandomUser() string {
	rand.Seed(time.Now().UTC().UnixNano())
	uc, _ := getSetting("USER_COUNT")
	n := int64(rand.Intn(uc))
	un, _ := Client.ZRange(Ctx, "username", n, n).Result()
	return un[0]
}

func ExistsUser(username string) bool {
	return ExistsValInZSET(username, "username")
}

func InsertUsers(UsersStruct *randomusers.UsersStruct) {
	for i := 0; i < len(UsersStruct.Results); i++ {
		if !ExistsUser(UsersStruct.Results[i].Login.Username) {
			id, _ := getSetting("USER_COUNT")
			Client.Incr(Ctx, "USER_COUNT")
			Client.ZAdd(Ctx, "username", &redis.Z{
				Score:  float64(id),
				Member: UsersStruct.Results[i].Login.Username,
			})
			Client.RPush(Ctx, "firstname", UsersStruct.Results[i].Name.First)
			Client.RPush(Ctx, "lastname", UsersStruct.Results[i].Name.Last)
			Client.ZAdd(Ctx, "email", &redis.Z{
				Score:  float64(id),
				Member: UsersStruct.Results[i].Email,
			})
			Client.RPush(Ctx, "regdate", UsersStruct.Results[i].Registered.Date)
		}
	}
}
