package redis

import (
	"testing"
)

func TestReturnURL(t *testing.T) {
	want := "[redis.ReturnURL func to be implemented]"
	msg := ReturnURL()
	if want != msg {
		t.Fatalf(`"[redis.ReturnURL func to be implemented]" = %q want match for %#q`, msg, want)
	}
}
