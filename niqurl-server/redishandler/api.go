package redishandler

import (
	"errors"
	"strconv"
)

func ViewAPI(url string) (map[string]interface{}, error) {
	z := map[string]interface{}{}
	x, _ := RDB.ZScan(Ctx, "shorturl", 0, url, 0).Val()

	// check if value exists
	if len(x) > 0 {
		i, err := strconv.Atoi(x[1])
		if err != nil {
			return z, err
		}
		z["shorturl"] = url
		z["longurl"] = RDB.ZRange(Ctx, "longurl", int64(i), int64(i)).Val()[0]
		z["user"] = RDB.ZRange(Ctx, "username", int64(i), int64(i)).Val()[0]
		return z, nil
	}
	return z, errors.New("shorturl not found")
}

func RedirectAPI(url string) (string, error) {
	x, _ := RDB.ZScan(Ctx, "shorturl", 0, url, 0).Val()

	// check if value exists
	if len(x) > 0 {
		i, err := strconv.Atoi(x[1])
		if err != nil {
			return "", err
		}
		url = RDB.ZRange(Ctx, "longurl", int64(i), int64(i)).Val()[0]
		return url, nil
	}
	return "", errors.New("shorturl not found")
}
