package cli

import (
	"math/rand"
	"time"
)

func MakeURL(longurl string) (string, string, error) {
	// redis.QueryURL to check if longurl was converted before
	// if err != nil {
	// redis.ReturnURL(longurl)
	// }

	b := make([]byte, SHORT_URL_LEN)

	// for loop with redis.QueryURL to check if this shorturl already exists
	//
	// for redis.QueryURL(b) {
	//		for i := range b {
	//		b[i] = ALLOWED_URL_CHARS[rand.Intn(len(ALLOWED_URL_CHARS))]
	//     }
	//  }
	// return longurl, string(b), nil

	// loop for generating random url
	for i := range b {
		rand.Seed(time.Now().UTC().UnixNano())
		b[i] = ALLOWED_URL_CHARS[rand.Intn(len(ALLOWED_URL_CHARS))]
	}
	return longurl, string(b), nil

}
