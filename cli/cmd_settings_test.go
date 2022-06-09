package cli

import (
	r "redishandler"
	"testing"
)

func TestShowSettingsNoErr(t *testing.T) {
	err := cmdSettingsGeneric()
	if err != nil {
		t.Fatalf(`Error: %q want match for nil`, err.Error())
	}
}

func TestShowSettingsErr(t *testing.T) {
	uc, _ := r.RDB.Get(r.Ctx, "USER_WAIT_TIME").Int()
	r.RDB.Del(r.Ctx, "USER_WAIT_TIME")
	err := cmdSettingsGeneric()
	r.RDB.Set(r.Ctx, "USER_WAIT_TIME", uc, 0)
	if err == nil {
		t.Fatal(`Error: nil want match for err`)
	}
}
