package redishandler

import (
	"github.com/go-redis/redis/v8"
)

//
func SearchURL(url, ver string, RDB *redis.Client) bool {
	// s can be "[val [url val]]"" or "[val []]""
	s := RDB.Do(Ctx, "zscan", ver, "0", "match", url).String()
	s = s[len(s)-3:]
	if s == "[]]" {
		return false
	}
	return true
}
