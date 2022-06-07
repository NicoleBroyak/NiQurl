package cli

import (
	"redishandler"
	"testing"
)

func TestLoadSettingsNoErr(t *testing.T) {
	RDB := redishandler.RedisStart()
	defer RDB.Close()
	err := LoadSettings()
	if err != nil {
		t.Fatalf(`Error: %q want match for nil`, err.Error())
	}
}

func TestLoadSettingsErr(t *testing.T) {
	RDB := redishandler.RedisStart()
	defer RDB.Close()
	uc, _ := RDB.Get(redishandler.Ctx, "USER_WAIT_TIME").Int()
	RDB.Del(redishandler.Ctx, "USER_WAIT_TIME")
	err := LoadSettings()
	RDB.Set(redishandler.Ctx, "USER_WAIT_TIME", uc, 0)
	if err == nil {
		t.Fatal(`Error: nil want match for err`)
	}
}
