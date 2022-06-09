package cli

import (
	"testing"
)

func TestMakeURLNoErr(t *testing.T) {
	err := cmdMakeURLGeneric("http://google.com")
	if err != nil {
		t.Fatalf(`Error: %q want match for nil`, err.Error())
	}
}

func FuzzMakeURL(f *testing.F) {
	testcases := []string{"https://google.com", "http://asf.pl", "test.waw.pl"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, url string) {
		err := cmdMakeURLGeneric(url)
		if err != nil {
			t.Errorf(`Error: %q want match for nil`, err.Error())
		}
	})
}
