package urlhandler

import "testing"

func TestNiqURLString(t *testing.T) {
	_, err := StringToNiqURL("")
	if err == nil {
		t.Fatalf(`Error: %v want match for error`, err)
	}
	NiqURL, err := StringToNiqURL("google.com")
	if err != nil {
		t.Fatalf(`Error: %v want match for nil`, err)
	}
	if NiqURL.LongURL != "google.com" {
		t.Fatalf(`Error: %q want match for "https://google.com"`, "google.com")
	}
	NiqURL.IfEmptySchemeAddHTTPS()
	if NiqURL.LongURL != "https://google.com" {
		t.Fatalf(`Error: %q want match for "https://google.com"`, "https://google.com")
	}
	NiqURL.GenerateShortURLPath(5)
	if len(NiqURL.ShortURL) != 5 {
		t.Fatalf(`Error: %v want match for "5"`, len(NiqURL.ShortURL))
	}

}
