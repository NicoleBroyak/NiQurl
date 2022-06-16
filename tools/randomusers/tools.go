package randomusers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type UsersStruct struct {
	Results []struct {
		Name struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Email string `json:"email"`
		Login struct {
			Username string `json:"username"`
		} `json:"login"`
		Registered struct {
			Date time.Time `json:"date"`
		} `json:"registered"`
	} `json:"results"`
}

func QueryRandomUsersAPI(number int) ([]byte, error) {
	apiSource := fmt.Sprintf("https://randomuser.me/api/?results=%v&inc=login,name,email,registered", number)
	response, err := http.Get(apiSource)
	if err != nil {
		return []byte{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
func GenerateFakeUsers(num int) (*UsersStruct, error) {
	var UsersStruct *UsersStruct
	QueryJSON, err := QueryRandomUsersAPI(num)
	if err != nil {
		return UsersStruct, err
	}
	UsersStruct, err = UsersStruct.New(QueryJSON)
	if err != nil {
		log.Print("Error related with random user API. Try again later")
		return UsersStruct, err
	}
	log.Print("Generating random users...")
	return UsersStruct, err
}

// returns UsersStruct with specified number of users
func (User UsersStruct) New(QueryJSON []byte) (*UsersStruct, error) {
	err := json.Unmarshal(QueryJSON, &User)
	if err != nil {
		return &UsersStruct{}, err
	}

	if len(User.Results) == 0 {
		return &UsersStruct{}, errors.New("error with getting users")

	}
	return &User, nil
}
