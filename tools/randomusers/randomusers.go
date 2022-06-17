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
	queryJSON, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return queryJSON, nil
}

func GenerateFakeUsers(num int) *UsersStruct {
	queryJSON, err := QueryRandomUsersAPI(num)
	if err != nil {
		panic("couldn't generate random users, aborting app, try again later")
	}
	usersStruct, err := UsersStruct{}.NewFromAPI(queryJSON)
	if err != nil {
		panic("couldn't generate random users, aborting app, try again later")
	}
	return usersStruct
}

// returns UsersStruct from API Query result
func (usersStruct UsersStruct) NewFromAPI(QueryJSON []byte) (*UsersStruct, error) {
	err := json.Unmarshal(QueryJSON, &usersStruct)
	if err != nil {
		log.Print("Error related with random user API. Try again later")
		return &UsersStruct{}, err
	}

	if len(usersStruct.Results) == 0 {
		return &UsersStruct{}, errors.New("error with getting users")

	}
	log.Print("Generating random users...")
	return &usersStruct, nil
}
