package redis

import (
	"testing"
)

func TestQueryURL(t *testing.T) {
	want := "[redis.QueryURL func to be implemented]"
	msg := QueryURL()
	if want != msg {
		t.Fatalf(`"[redis.QueryURL func to be implemented]" = %q want match for %#q`, msg, want)
	}
}
