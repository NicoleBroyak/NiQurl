package cli

import (
	"errors"
	r "redishandler"
)

var SHORT_URL_LEN int
var USER_WAIT_TIME int
var USER_COUNT int
var URL_COUNT int

const ALLOWED_URL_CHARS string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LoadSettings() error {
	RDB := r.RedisStart()
	defer RDB.Close()

	a := map[string]*int{
		"SHORT_URL_LEN":  &SHORT_URL_LEN,
		"USER_WAIT_TIME": &USER_WAIT_TIME,
		"USER_COUNT":     &USER_COUNT,
		"URL_COUNT":      &URL_COUNT}

	for k, v := range a {
		b, _ := RDB.Get(r.Ctx, k).Result()
		if len(b) == 0 {
			return errors.New("Value " + k + " not found")
		}
		x, _ := RDB.Get(r.Ctx, k).Int()
		*v = x
	}
	return nil
}
