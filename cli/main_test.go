package cli

import (
	r "redishandler"
	"testing"
)

func TestMain(m *testing.M) {
	r.CheckVariables()
	m.Run()
	r.RDB.FlushAll(r.Ctx)
}
