# randomusers
--
    import "github.com/nicolebroyak/niqurl/tools/randomusers"


## Usage

#### func  QueryRandomUsersAPI

```go
func QueryRandomUsersAPI(number int) ([]byte, error)
```

#### type UsersStruct

```go
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
```


#### func (UsersStruct) NewFromAPI

```go
func (User UsersStruct) NewFromAPI(QueryJSON []byte) (*UsersStruct, error)
```
returns UsersStruct from API Query result
