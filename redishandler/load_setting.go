package redishandler

import (
	"errors"
)

// "SHORT_URL_LEN", "USER_WAIT_TIME", "USER_COUNT" or "URL_COUNT"
func loadSetting(setting string) (int, error) {

	s, err := RDB.Get(Ctx, setting).Result()
	if err != nil {
		return 0, err
	}
	if len(s) == 0 {
		return 0, errors.New("Value " + setting + " not found")
	}
	v, _ := RDB.Get(Ctx, setting).Int()
	return v, nil
}
