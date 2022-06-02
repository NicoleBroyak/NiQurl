package redis

import "fmt"

func WaitTime(user string) string {
	rdb := Start()
	u, _ := rdb.Get(ctx, "user_count").Result()
	id := "url:" + u
	fmt.Println(id)
	return ""
}
