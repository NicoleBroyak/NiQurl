package redishandler

import (
	"testing"
)

func TestLoadSettingsNoErr(t *testing.T) {
	_, err := LoadSettings()
	if err != nil {
		t.Fatalf(`Error: %q want match for nil`, err.Error())
	}
}

func TestLoadSettingsErr(t *testing.T) {
	u, _ := RDB.Get(Ctx, "USER_WAIT_TIME").Int()
	RDB.Del(Ctx, "USER_WAIT_TIME")
	_, err := LoadSettings()
	RDB.Set(Ctx, "USER_WAIT_TIME", u, 0)
	if err == nil {
		t.Fatal(`Error: nil want match for err`)
	}
}
