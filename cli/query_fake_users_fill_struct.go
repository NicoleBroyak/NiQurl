package cli

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func QFUFillStruct(url string, Users *UsersStruct) [4]error {
	res, err := http.Get(url)
	body, err2 := io.ReadAll(res.Body)
	err3 := json.Unmarshal(body, &Users)
	err4 := err3
	if len(Users.Results) == 0 {
		err4 = errors.New("Error with getting users")
	}
	return [4]error{err, err2, err3, err4}
}
