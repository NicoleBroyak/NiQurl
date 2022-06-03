package redis

import (
	"testing"
)

func TestAddUsers(t *testing.T) {
	want := "aa"
	msg := AddUsers("a")
	if want != msg {
		t.Fatalf(`[aa] = %q want match for %#q, nil`, msg, want)
	}
}

func FuzzAddUsers(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		double := AddUsers(orig)
		want := orig + orig
		if double != want {
			t.Errorf("Before: %q, after: %q", orig, double)
		}
	})
}
