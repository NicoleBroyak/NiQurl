package redishandler

func InsertUser(userdata [5]string) {
	RDB := RedisStart()
	defer RDB.Close()

	id, _ := RDB.Get(Ctx, "USER_COUNT").Int()
	RDB.Do(Ctx, "incr", "USER_COUNT")
	RDB.Do(Ctx, "ZADD", "username", id, userdata[0])
	RDB.Do(Ctx, "RPUSH", "firstname", userdata[1])
	RDB.Do(Ctx, "RPUSH", "lastname", userdata[2])
	RDB.Do(Ctx, "ZADD", "email", id, userdata[3])
	RDB.Do(Ctx, "RPUSH", "regdate", userdata[4])
}
