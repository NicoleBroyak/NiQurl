package redishandler

import (
	"fmt"
	"math/rand"

	"time"
)

func AssignRandUser() (string, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	uc, err := RDB.Get(Ctx, "USER_COUNT").Int()
	if err != nil {
		return "", err
	}
	randnum := int64(rand.Intn(uc))
	user, _ := RDB.ZRange(Ctx, "username", randnum, randnum).Result()
	if len(user) != 0 {
		return user[0], nil
	}
	return "", fmt.Errorf("Empty user, %v", randnum)

}
