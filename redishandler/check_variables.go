package redishandler

import (
	"log"
)

func CheckVariables() {

	_, err := loadSetting("SHORT_URL_LEN")
	_, err2 := loadSetting("USER_WAIT_TIME")
	_, err3 := loadSetting("URL_COUNT")
	_, err4 := loadSetting("USER_COUNT")

	if err != nil || err2 != nil {
		RDB.Set(Ctx, "USER_WAIT_TIME", 30, 0)
		RDB.Set(Ctx, "SHORT_URL_LEN", 4, 0)
		log.Println("init vars USER_WAIT_TIME and SHORT_URL_LEN" +
			"not found in redis, initializing default database keys")
	}
	if err3 != nil {
		RDB.Set(Ctx, "URL_COUNT", 0, 0)
		log.Println("init var URL_COUNT not found in redis, " +
			"initializing database key")
	}
	if err4 != nil {
		RDB.Set(Ctx, "USER_COUNT", 0, 0)
		generateFakeUsersGeneric(5, 1, 1000)
		log.Println("init var USER not found in redis, " +
			"initializing database key and " +
			"generating 5 random users" +
			"(see -generate-fake-users flag)")
	}
}
