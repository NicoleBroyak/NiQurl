package cli

import "testing"

func TestSetTimeErr(t *testing.T) {
	err := setTimeFunc(0)
	if err == nil {
		t.Fatal(`Error: nil want match for err`)
	}
}

func TestSetTimeNoErr(t *testing.T) {
	err := setTimeFunc(10)
	if err != nil {
		t.Fatalf(`Error: %q want match for nil`, err.Error())
	}
}

func FuzzSetTime(f *testing.F) {
	testcases := []int{-1000, -500, -211, 0, 15, 256}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, time int) {
		err := setTimeFunc(time)
		if time > 0 {
			if err != nil {
				t.Fatalf(`Error: %q want match for nil`, err.Error())
			}
		}
		if time <= 0 {
			if err == nil {
				t.Fatal(`Error: nil want match for err`)
			}
		}
	})
}
