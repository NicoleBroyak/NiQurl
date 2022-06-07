package cli

import "testing"

func TestSetLenErr(t *testing.T) {
	err := setLenFunc(0)
	if err == nil {
		t.Fatal(`Error: nil want match for err`)
	}
}

func TestSetLenNoErr(t *testing.T) {
	err := setLenFunc(10)
	if err != nil {
		t.Fatalf(`Error: %q want match for nil`, err.Error())
	}
}

func FuzzSetLen(f *testing.F) {
	testcases := []int{-1000, -500, -211, 0, 15, 256}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, len int) {
		err := setLenFunc(len)
		if len > 0 {
			if err != nil {
				t.Fatalf(`Error: %q want match for nil`, err.Error())
			}
		}
		if len <= 0 {
			if err == nil {
				t.Fatal(`Error: nil want match for err`)
			}
		}
	})
}
