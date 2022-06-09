package redishandler

import (
	"testing"
)

func TestGenerateFakeUsersPos(t *testing.T) {
	err := generateFakeUsersGeneric(1000, 1, 1000)

	// Skip test with error related to randomuser.me not app
	if err != nil && err.Error() == "invalid character '<' looking for beginning of value" {
		t.SkipNow()
	}
	if err != nil && err.Error() == "Error with getting users" {
		t.SkipNow()
	}

	if err != nil {
		t.Fatalf(`Error: %q want match for nil`, err.Error())
	}
}

func FuzzGenerateFakeUsers(f *testing.F) {
	testcases := []int{-100, -50, -1, 0, 20, 233, 3000}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, num int) {
		err := generateFakeUsersGeneric(num, 1, 1000)

		// Skip test with error related to randomuser.me not app
		if err != nil && err.Error() == "invalid character '<' looking for beginning of value" {
			t.SkipNow()
		}
		if err != nil && err.Error() == "Error with getting users" {
			t.SkipNow()
		}

		if err != nil && (num >= 1 && num <= 1000) {
			t.Errorf(`Error: %q want match for nil`, err.Error())
		}
		if err == nil && !(num >= 1 && num <= 1000) {
			t.Errorf(`Error: nil match for err`)
		}
	})
}
