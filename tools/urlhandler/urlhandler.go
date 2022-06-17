package urlhandler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"time"
)

// Alias for *url.URL type
type NiqURL struct {
	*url.URL
	LongURL  string
	ShortURL string
	UserName string
}

func StringToNiqURL(input string) (*NiqURL, error) {
	genericURL, err := url.Parse(input)
	NiqURL := &NiqURL{genericURL, genericURL.String(), "", ""}
	if err != nil {
		return NiqURL, err
	}
	return NiqURL, nil
}

func (NiqURL *NiqURL) IfEmptySchemeAddHTTPS() *NiqURL {
	if !NiqURL.IsAbs() {
		NiqURL.Scheme = "https"
		NiqURL.LongURL = NiqURL.String()
		log.Println("Adding https:// to non absoulute url...")
	}
	return NiqURL
}

func (NiqURL *NiqURL) GenerateShortURLPath(urlLen int) {
	rand.Seed(time.Now().UnixNano())
	longurl := NiqURL.String()
	longurlbyte := []byte(longurl)
	urlAsMD5 := md5.Sum(longurlbyte)
	stringMD5 := fmt.Sprintf("%x", urlAsMD5)
	randIndex := rand.Intn(31 - urlLen)
	NiqURL.ShortURL = stringMD5[randIndex : randIndex+urlLen]
}

func ShortenURLError() error {
	err := errors.New("can't generate short url with" +
		"specified length, please change" +
		"length using setlen command")
	fmt.Println(err)
	return err
}
