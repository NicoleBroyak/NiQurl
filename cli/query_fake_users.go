package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"redishandler"
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

func QueryFakeUsers(i int) {
	Users := UsersStruct{}
	fmt.Println("Generating random users...")
	url := fmt.Sprintf("https://randomuser.me/api/?results=%d&inc=login,name,email,registered", i+1)
	res, _ := http.Get(url)
	body, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(body, &Users)
	if err == nil {
		for index := 0; index < i; index++ {
			username := Users.Results[index].Login.Username
			firstname := Users.Results[index].Name.First
			lastname := Users.Results[index].Name.Last
			email := Users.Results[index].Email
			regdate := Users.Results[index].Registered.Date
			userdata := [4]string{username, firstname, lastname, email}
			redishandler.InsertUser(userdata, regdate)
			USER_COUNT += 1
		}
	}

}
