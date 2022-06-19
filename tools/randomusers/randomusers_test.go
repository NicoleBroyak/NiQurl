package randomusers

import (
	"testing"

	"github.com/nicolebroyak/niqurl/config/niqurlconfigs"
)

func TestRandomUsers(t *testing.T) {
	apiSource := "google.com"
	queryJSON, err := QueryRandomUsersAPI(apiSource, 5)
	_, err2 := UsersStruct{}.NewFromAPI(queryJSON)
	if err == nil {
		t.Fatalf(`Error: %v want match for nil`, err)
	}
	if err2 == nil {
		t.Fatalf(`Error: %v want match for nil`, err2)
	}
	apiSource = niqurlconfigs.CreateAPISourceFromDefault(10)
	usersStruct := GenerateFakeUsers(apiSource, 10)
	if len(usersStruct.Results) != 10 {
		t.Fatalf(`Error: %v want match for 10`, len(usersStruct.Results))
	}

}
