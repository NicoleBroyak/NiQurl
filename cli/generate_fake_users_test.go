package cli

import (
	"testing"
)

func TestGenerateFakeUsersPos(t *testing.T) {
	err := generateFakeUsers(1000)

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

func TestGenerateFakeUsersTooHigh(t *testing.T) {
	err := generateFakeUsers(22000)
	if err == nil {
		t.Fatalf(`Error: nil want match for error`)
	}
}

func TestGenerateFakeUsersNeg(t *testing.T) {
	err := generateFakeUsers(-3000)
	if err == nil {
		t.Fatalf(`Error: nil want match for error`)
	}
}

func FuzzGenerateFakeUsers(f *testing.F) {
	testcases := []int{-100, -50, -1, 0, 20, 233, 3000}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, fakeusers int) {
		err := generateFakeUsers(fakeusers)

		// Skip test with error related to randomuser.me not app
		if err != nil && err.Error() == "invalid character '<' looking for beginning of value" {
			t.SkipNow()
		}
		if err != nil && err.Error() == "Error with getting users" {
			t.SkipNow()
		}

		// verify if func raises error when not needed
		if err != nil && fakeusers > 0 && fakeusers <= 1000 {
			t.Errorf(`Error: %q want match for nil`, err.Error())
		}

		// verify if func raises error when needed
		if err == nil && fakeusers <= 0 {
			t.Error(`Error: nil want match for error`)
		}
		if err == nil && fakeusers > 1000 {
			t.Error(`Error: nil want match for error`)
		}
	})
}
