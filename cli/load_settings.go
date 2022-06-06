package cli

import (
	"redishandler"

	"github.com/go-redis/redis/v8"
)

var SHORT_URL_LEN int
var USER_WAIT_TIME int
var USER_COUNT int
var URL_COUNT int

const ALLOWED_URL_CHARS string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LoadSettings(RDB *redis.Client) error {

	a := map[string]*int{
		"SHORT_URL_LEN":  &SHORT_URL_LEN,
		"USER_WAIT_TIME": &USER_WAIT_TIME,
		"USER_COUNT":     &USER_COUNT,
		"URL_COUNT":      &URL_COUNT}

	for k, v := range a {
		x, err := RDB.Get(redishandler.Ctx, k).Int()
		if err != nil {
			return err
		}
		*v = x
	}
	return nil
}
