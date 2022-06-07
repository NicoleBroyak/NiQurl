package cli

import (
	"errors"
	r "redishandler"
)

func SetVariables() error {
	RDB := r.RedisStart()
	defer RDB.Close()

	_, err := RDB.Get(r.Ctx, "SHORT_URL_LEN").Result()
	_, err2 := RDB.Get(r.Ctx, "USER_WAIT_TIME").Result()

	if err != nil || err2 != nil {
		RDB.Set(r.Ctx, "USER_WAIT_TIME", 30, 0)
		RDB.Set(r.Ctx, "SHORT_URL_LEN", 4, 0)
		i, err3 := r.CheckIntVar("URL_COUNT", RDB)
		if err3 != nil {
			return err3
		}
		j, err4 := r.CheckIntVar("USER_COUNT", RDB)
		if err4 != nil {
			return err4
		}
		RDB.Set(r.Ctx, "URL_COUNT", i, 0)
		RDB.Set(r.Ctx, "USER_COUNT", j, 0)
		LoadSettings()
		return errors.New("init vars not found in redis, initializing database keys")
	}
	return nil
}
