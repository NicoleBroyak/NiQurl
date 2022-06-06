package cli

import (
	"fmt"
	"redishandler"
)

func SetVariables() error {
	RDB := redishandler.RedisStart()
	defer RDB.Close()

	_, err := RDB.Get(redishandler.Ctx, "SHORT_URL_LEN").Int()
	_, err2 := RDB.Get(redishandler.Ctx, "USER_WAIT_TIME").Int()

	if err != nil || err2 != nil {
		RDB.Set(redishandler.Ctx, "USER_WAIT_TIME", 30, 0)
		RDB.Set(redishandler.Ctx, "SHORT_URL_LEN", 4, 0)
		i, err := redishandler.CheckIntVar("URL_COUNT", RDB)
		if err != nil {
			return err
		}
		j, err := redishandler.CheckIntVar("USER_COUNT", RDB)
		if err != nil {
			return err
		}
		RDB.Set(redishandler.Ctx, "URL_COUNT", i, 0)
		RDB.Set(redishandler.Ctx, "USER_COUNT", j, 0)
		LoadSettings(RDB)
		fmt.Println("initializing database keys")
	}
	return nil
}
