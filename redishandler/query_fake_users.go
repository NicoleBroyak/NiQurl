package redishandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func QueryFakeUsers(i int) error {
	Users := UsersStruct{}
	fmt.Println("Generating random users...")
	url := fmt.Sprintf("https://randomuser.me/api/?results=%d&inc=login,name,email,registered", i+1)

	errs := qfuFillStruct(url, &Users)
	err := qfuErrHandler(errs)

	if err == nil {
		for index := 0; index < i; index++ {
			qfuCreateUser(&Users, index)
		}
		return nil
	}
	return err
}

func qfuCreateUser(Users *UsersStruct, index int) {
	username := Users.Results[index].Login.Username
	firstname := Users.Results[index].Name.First
	lastname := Users.Results[index].Name.Last
	email := Users.Results[index].Email
	regdate := Users.Results[index].Registered.Date.String()
	userdata := [5]string{username, firstname, lastname, email, regdate}
	insertUser(userdata)
}

func qfuErrHandler(errs [4]error) error {
	for _, v := range errs {
		if v != nil {
			return v
		}
	}
	return nil

}

func qfuFillStruct(url string, Users *UsersStruct) [4]error {
	res, err := http.Get(url)
	body, err2 := io.ReadAll(res.Body)
	err3 := json.Unmarshal(body, &Users)
	err4 := err3
	if len(Users.Results) == 0 {
		err4 = errors.New("Error with getting users")
	}
	return [4]error{err, err2, err3, err4}
}

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
