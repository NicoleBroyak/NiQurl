package cli

import (
	"testing"
)

func TestCLIInitializeNoErr(t *testing.T) {
	err := cliInitialize(2)
	if err != nil {
		t.Fatalf(`Error: %q want match for nil`, err.Error())
	}
}

func TestCLIInitializeErr(t *testing.T) {
	err := cliInitialize(-10)
	if err == nil {
		t.Fatalf(`Error: nil want match for err`)
	}
}

func FuzzCLIInitialize(f *testing.F) {
	testcases := []int{-1000, -500, -211, 0, 15, 256}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, uc int) {
		err := cliInitialize(uc)
		if uc > 0 {
			if err != nil {
				t.Fatalf(`Error: %q want match for nil`, err.Error())
			}
		}
		if uc <= 0 {
			if err == nil {
				t.Fatal(`Error: nil want match for err`)
			}
		}
	})
}
