package cli

import "redishandler"

func QFUCreateUser(Users *UsersStruct, index int) {
	username := Users.Results[index].Login.Username
	firstname := Users.Results[index].Name.First
	lastname := Users.Results[index].Name.Last
	email := Users.Results[index].Email
	regdate := Users.Results[index].Registered.Date.String()
	userdata := [5]string{username, firstname, lastname, email, regdate}
	redishandler.InsertUser(userdata)
	USER_COUNT += 1
}
