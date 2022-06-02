package redis

import "fmt"

func MakeURL(longurl, shorturl, user string) {
	rdb := Start()
	u, _ := rdb.Get(ctx, "user_count").Result()
	id := "url:" + u
	fmt.Println(id)
	rdb.Do(ctx, "incr", "user_count")
	rdb.Do(ctx, "HMSET", id, "longurl", longurl, "shorturl", shorturl, "createdby", user)
}
