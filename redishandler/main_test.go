package redishandler

import (
	"testing"
)

func TestMain(m *testing.M) {
	CheckVariables()
	m.Run()
	RDB.FlushAll(Ctx)
}
