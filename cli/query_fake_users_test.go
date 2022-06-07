package cli

import (
	"testing"
)

func TestQueryFakeUsersPos(t *testing.T) {
	err := QueryFakeUsers(10)
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

func TestQueryFakeUsersTooHigh(t *testing.T) {
	err := QueryFakeUsers(10000)
	if err == nil {
		t.Fatalf(`Error: nil want match for error`)
	}
}

func TestQueryFakeUsersTooLow(t *testing.T) {
	err := QueryFakeUsers(-20)
	if err == nil {
		t.Fatalf(`Error: nil want match for error`)
	}
}

func FuzzQueryFakeUsers(f *testing.F) {
	testcases := []int{1, 3, 5, 9}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, fakeusers int) {
		err := QueryFakeUsers(fakeusers)
		if err != nil && err.Error() == "invalid character '<' looking for beginning of value" {
			t.SkipNow()
		}
		if err != nil && err.Error() == "Error with getting users" {
			t.SkipNow()
		}
		if err != nil && (fakeusers <= 1000 && fakeusers >= 1) {
			t.Errorf(`Error: %q want match for nil || fakeusers: %v`, err.Error(), fakeusers)
		}
		if err == nil && (fakeusers > 1000 || fakeusers < 1) {
			t.Error(`Error: nil want match for error`)
		}
	})
}
