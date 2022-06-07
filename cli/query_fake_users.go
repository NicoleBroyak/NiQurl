package cli

import (
	"errors"
	"fmt"
	"time"
)

type UsersStruct struct {
	Results []struct {
		Name struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Email string `json:"email"`
		Login struct {
			UUID     string `json:"uuid"`
			Username string `json:"username"`
			Password string `json:"password"`
			Salt     string `json:"salt"`
			Md5      string `json:"md5"`
			Sha1     string `json:"sha1"`
			Sha256   string `json:"sha256"`
		} `json:"login"`
		Registered struct {
			Date time.Time `json:"date"`
			Age  int       `json:"age"`
		} `json:"registered"`
	} `json:"results"`
	Info struct {
		Seed    string `json:"seed"`
		Results int    `json:"results"`
		Page    int    `json:"page"`
		Version string `json:"version"`
	} `json:"info"`
}

func QueryFakeUsers(i int) error {
	if i > 1000 || i < 1 {
		return errors.New("You can generate between 1 and 1000 users one time")
	}

	Users := UsersStruct{}
	fmt.Println("Generating random users...")
	url := fmt.Sprintf("https://randomuser.me/api/?results=%d&inc=login,name,email,registered", i+1)

	errs := QFUFillStruct(url, &Users)
	err := QFUErrHandler(errs)

	if err == nil {
		for index := 0; index < i; index++ {
			QFUCreateUser(&Users, index)
		}
		return nil
	}
	return err
}
