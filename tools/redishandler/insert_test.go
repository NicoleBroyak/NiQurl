package redishandler

import (
	"testing"
	"time"

	"github.com/nicolebroyak/niqurl/tools/randomusers"
	"github.com/nicolebroyak/niqurl/tools/urlhandler"
)

func TestInsertIntoSortedSet(t *testing.T) {
	insertIntoSortedSet(
		"username",
		float64(client.ZCount(context, "username", "-inf", "+inf").Val()+1),
		"NewUserName23Randname",
	)
	newUserQuery, _ := client.ZScan(context, "username", 0, "NewUserName23Randname", 0).Val()
	client.ZPopMax(context, "username")
	if newUserQuery[0] != "NewUserName23Randname" {
		t.Fatalf(`error: %v want match for %v`, newUserQuery[0], "NewUserName23Randname")
	}
}

func TestInsertIntoList(t *testing.T) {
	insertIntoList("createdby", "NewUserName23Randname")
	newUserQuery := client.LRange(context, "createdby", -1, -1).Val()
	client.RPop(context, "createdby")
	if newUserQuery[0] != "NewUserName23Randname" {
		t.Fatalf(`error: %q want match for %q`, newUserQuery[0], "NewUserName23Randname")
	}
}

func TestInsertURLAuthor(t *testing.T) {
	insertURLAuthor("NewUserName23Randname")
	newUserQuery := client.LRange(context, "createdby", -1, -1).Val()
	client.RPop(context, "createdby")
	if newUserQuery[0] != "NewUserName23Randname" {
		t.Fatalf(`error: %q want match for %q`, newUserQuery[0], "NewUserName23Randname")
	}
}

func TestInsertFirstName(t *testing.T) {
	insertFirstName("Firstname")
	firstNameQuery := client.LRange(context, "firstname", -1, -1).Val()
	client.RPop(context, "firstname")
	if firstNameQuery[0] != "Firstname" {
		t.Fatalf(`error: %q want match for %q`, firstNameQuery[0], "Firstname")
	}
}

func TestInsertLastName(t *testing.T) {
	insertLastName("Lastname")
	lastNameQuery := client.LRange(context, "lastname", -1, -1).Val()
	client.RPop(context, "lastname")
	if lastNameQuery[0] != "Lastname" {
		t.Fatalf(`error: %q want match for %q`, lastNameQuery[0], "Firstname")
	}
}

func TestInsertEmail(t *testing.T) {
	email := "newrandomemail@gmail.com"
	insertEmail(email)
	emailQuery, _ := client.ZScan(context, "email", 0, email, 0).Val()
	client.ZPopMax(context, "email")
	if emailQuery[0] != email {
		t.Fatalf(`error: %v want match for %v`, emailQuery[0], email)
	}
}

func TestInsertLongURL(t *testing.T) {
	longURL := "https://newrandomwebsitename.com/random/random2"
	insertLongURL(longURL)
	longURLQuery, _ := client.ZScan(context, "longurl", 0, longURL, 0).Val()
	client.ZPopMax(context, "longurl")
	if longURLQuery[0] != longURL {
		t.Fatalf(`error: %v want match for %v`, longURLQuery[0], longURL)
	}
}

func TestInsertShortURL(t *testing.T) {
	shortURL := "zx224ag"
	insertShortURL(shortURL)
	shortURLQuery, _ := client.ZScan(context, "shorturl", 0, shortURL, 0).Val()
	client.ZPopMax(context, "shorturl")
	if shortURLQuery[0] != shortURL {
		t.Fatalf(`error: %v want match for %v`, shortURLQuery[0], shortURL)
	}
}

func TestInsertUser(t *testing.T) {
	username := "NewUserName23Randname"
	insertUserName(username)
	usernameQuery, _ := client.ZScan(context, "username", 0, username, 0).Val()
	client.ZPopMax(context, "username")
	if usernameQuery[0] != username {
		t.Fatalf(`error: %v want match for %v`, usernameQuery[0], username)
	}
}

func TestInsertUserRegDate(t *testing.T) {
	regdate := time.Now()
	insertRegistrationDate(regdate)
	regdateQuery := client.LRange(context, "regdate", -1, -1).Val()
	client.RPop(context, "regdate")
	if regdateQuery[0] != regdate.String() {
		t.Fatalf(`error: %q want match for %q`, regdateQuery[0], regdate.String())
	}
}

func TestInsertUserData(t *testing.T) {
	firstname := "Firstname"
	lastname := "Lastname"
	email := "newrandomemail@gmail.com"
	username := "NewUserName23Randname"
	regdate := time.Now()
	userstruct := randomusers.UsersStruct{}
	appendUserStruct(&userstruct)
	userstruct.Results[0].Email = email
	userstruct.Results[0].Registered.Date = regdate
	userstruct.Results[0].Login.Username = username
	userstruct.Results[0].Name.First = firstname
	userstruct.Results[0].Name.Last = lastname
	insertUserData(&userstruct, 0)
	firstNameQuery := client.LRange(context, "firstname", -1, -1).Val()
	lastNameQuery := client.LRange(context, "lastname", -1, -1).Val()
	emailQuery, _ := client.ZScan(context, "email", 0, email, 0).Val()
	usernameQuery, _ := client.ZScan(context, "username", 0, username, 0).Val()
	regdateQuery := client.LRange(context, "regdate", -1, -1).Val()
	client.RPop(context, "regdate")
	client.RPop(context, "firstname")
	client.RPop(context, "firstname")
	client.ZPopMax(context, "username")
	client.ZPopMax(context, "email")
	client.Decr(context, "USER_COUNT")

	if regdateQuery[0] != regdate.String() ||
		firstNameQuery[0] != firstname ||
		lastNameQuery[0] != lastname ||
		emailQuery[0] != email ||
		usernameQuery[0] != username {
		t.Fatalf(`error:
		Firstname: %q want match for %q
		Lastname: %q want match for %q
		Email: %q want match for %q
		Username: %q want match for %q
		regdate: %q want match for %q`,
			firstNameQuery[0], firstname,
			lastNameQuery[0], lastname,
			emailQuery[0], email,
			usernameQuery[0], username,
			regdateQuery[0], regdate.String())
	}
}

func TestInsertUsersData(t *testing.T) {
	firstname, firstname2 := "Firstname", "Firstname2"
	lastname, lastname2 := "Lastname", "Lastname2"
	email, email2 := "newrandomemail@gmail.com", "newrandomemail2@gmail.com"
	username, username2 := "NewUserName23Randname", "NewUserName24Randname"
	regdate, regdate2 := time.Now(), time.Now().Add(time.Duration(time.Now().Day()))
	userstruct := randomusers.UsersStruct{}
	appendUserStruct(&userstruct)
	userstruct.Results[0].Email = email
	userstruct.Results[0].Registered.Date = regdate
	userstruct.Results[0].Login.Username = username
	userstruct.Results[0].Name.First = firstname
	userstruct.Results[0].Name.Last = lastname
	appendUserStruct(&userstruct)
	userstruct.Results[1].Email = email2
	userstruct.Results[1].Registered.Date = regdate2
	userstruct.Results[1].Login.Username = username2
	userstruct.Results[1].Name.First = firstname2
	userstruct.Results[1].Name.Last = lastname2
	InsertUsers(&userstruct)
	firstNameQuery := client.LRange(context, "firstname", -2, -2).Val()
	lastNameQuery := client.LRange(context, "lastname", -2, -2).Val()
	emailQuery, _ := client.ZScan(context, "email", 0, email, 0).Val()
	usernameQuery, _ := client.ZScan(context, "username", 0, username, 0).Val()
	regdateQuery := client.LRange(context, "regdate", -2, -2).Val()
	firstName2Query := client.LRange(context, "firstname", -1, -1).Val()
	lastName2Query := client.LRange(context, "lastname", -1, -1).Val()
	email2Query, _ := client.ZScan(context, "email", 0, email2, 0).Val()
	username2Query, _ := client.ZScan(context, "username", 0, username2, 0).Val()
	regdate2Query := client.LRange(context, "regdate", -1, -1).Val()
	client.RPop(context, "regdate")
	client.RPop(context, "firstname")
	client.RPop(context, "lastname")
	client.RPop(context, "regdate")
	client.RPop(context, "firstname")
	client.RPop(context, "lastname")
	client.ZPopMax(context, "username", 2)
	client.ZPopMax(context, "email", 2)
	client.Decr(context, "USER_COUNT")
	client.Decr(context, "USER_COUNT")

	if regdateQuery[0] != regdate.String() ||
		firstNameQuery[0] != firstname ||
		lastNameQuery[0] != lastname ||
		emailQuery[0] != email ||
		usernameQuery[0] != username {
		t.Fatalf(`error:
		Firstname: %q want match for %q
		Lastname: %q want match for %q
		Email: %q want match for %q
		Username: %q want match for %q
		regdate: %q want match for %q`,
			firstNameQuery[0], firstname,
			lastNameQuery[0], lastname,
			emailQuery[0], email,
			usernameQuery[0], username,
			regdateQuery[0], regdate.String())
	}

	if regdate2Query[0] != regdate2.String() ||
		firstName2Query[0] != firstname2 ||
		lastName2Query[0] != lastname2 ||
		email2Query[0] != email2 ||
		username2Query[0] != username2 {
		t.Fatalf(`error:
		Firstname: %q want match for %q
		Lastname: %q want match for %q
		Email: %q want match for %q
		Username: %q want match for %q
		regdate: %q want match for %q`,
			firstName2Query[0], firstname2,
			lastName2Query[0], lastname2,
			email2Query[0], email2,
			username2Query[0], username2,
			regdate2Query[0], regdate2.String())
	}
}

func TestInsertURLData(t *testing.T) {
	NiqURL, _ := urlhandler.StringToNiqURL("https://newrandomwebsitename.com/random/random2")
	NiqURL.ShortURL = "zx224ag"
	NiqURL.UserName = "NewUserName23Randname"
	InsertURLData(NiqURL)
	PrintUserWaitTime(NiqURL.UserName)
	ProcessExistingURL(NiqURL.LongURL)
	shortURLQuery := client.ZRange(context, "shorturl", -1, -1).Val()
	longURLQuery := client.ZRange(context, "longurl", -1, -1).Val()
	createdbyQuery := client.LRange(context, "createdby", -1, -1).Val()
	client.RPop(context, "createdby")
	client.ZPopMax(context, "shorturl")
	client.ZPopMax(context, "longurl")

	if shortURLQuery[0] != NiqURL.ShortURL ||
		longURLQuery[0] != NiqURL.LongURL ||
		createdbyQuery[0] != NiqURL.UserName {
		t.Fatalf(`error:
		LongURL: %q want match for %q
		ShortURL: %q want match for %q
		Username: %q want match for %q`,
			shortURLQuery[0], NiqURL.ShortURL,
			longURLQuery[0], NiqURL.LongURL,
			createdbyQuery[0], NiqURL.UserName)
	}
}

func appendUserStruct(userstruct *randomusers.UsersStruct) {
	userstruct.Results = append(userstruct.Results, struct {
		Name struct {
			First string "json:\"first\""
			Last  string "json:\"last\""
		} "json:\"name\""
		Email string "json:\"email\""
		Login struct {
			Username string "json:\"username\""
		} "json:\"login\""
		Registered struct {
			Date time.Time "json:\"date\""
		} "json:\"registered\""
	}{})
}
