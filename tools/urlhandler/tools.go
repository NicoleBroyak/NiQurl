package urlhandler

import (
	"log"
	"net/url"
)

//
type NiqURL struct {
	url *url.URL
}

func InputStringToNiqURL(input string) (*NiqURL, error) {
	genericURL, err := url.Parse(input)
	URL := &NiqURL{url: genericURL}
	if err != nil {
		return URL, err
	}
	return URL, nil
}

func (NiqurlURL *NiqURL) AddHTTPSSchemeIfNonAbs() *NiqURL {
	if !NiqurlURL.url.IsAbs() {
		NiqurlURL.url.Scheme = "https"
		log.Println("Adding https:// to non absoulute url...")
	}
	return NiqurlURL
}

func (NiqURL *NiqURL) String() string {
	return NiqURL.url.String()
}
