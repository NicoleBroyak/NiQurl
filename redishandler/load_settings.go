package redishandler

import (
	"errors"
)

func LoadSettings() (map[string]int, error) {
	m := map[string]int{
		"SHORT_URL_LEN":  0,
		"USER_WAIT_TIME": 0,
		"USER_COUNT":     0,
		"URL_COUNT":      0}

	for k := range m {
		b, err := RDB.Get(Ctx, k).Result()
		if err != nil {
			return map[string]int{}, err
		}
		if len(b) == 0 {
			return map[string]int{}, errors.New("Value " + k + " not found")
		}
		v, _ := RDB.Get(Ctx, k).Int()
		m[k] = v
	}
	return m, nil
}
