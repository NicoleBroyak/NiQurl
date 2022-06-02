package redis

import (
	"testing"
)

func TestInit(t *testing.T) {
	want := "[redis.Init func to be implemented]"
	msg := Init()
	if want != msg {
		t.Fatalf(`"[redis.Init func to be implemented]" = %q want match for %#q, nil`, msg, want)
	}
}
