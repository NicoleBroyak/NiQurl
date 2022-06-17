# redishandler
--
    import "github.com/nicolebroyak/niqurl/tools/redishandler"


## Usage

#### func  ChangeSetting

```go
func ChangeSetting(setting string, value int)
```

#### func  ExistsLongURL

```go
func ExistsLongURL(longurl string) bool
```

#### func  ExistsShortURL

```go
func ExistsShortURL(shorturl string) bool
```

#### func  GetIndexOfShortURL

```go
func GetIndexOfShortURL(shorturl string) (int64, error)
```

#### func  GetLongURL

```go
func GetLongURL(index int64) string
```

#### func  GetRandomUser

```go
func GetRandomUser() string
```

#### func  GetSetting

```go
func GetSetting(setting string) int
```
function assumes that setting was validated before invoking

#### func  GetUser

```go
func GetUser(index int64) string
```

#### func  InsertURLData

```go
func InsertURLData(NiqURL urlhandler.NiqURL)
```

#### func  InsertUsers

```go
func InsertUsers(UsersStruct *randomusers.UsersStruct)
```

#### func  IsUserOnWaitTime

```go
func IsUserOnWaitTime(user string) bool
```

#### func  PrintCurrentCLISettings

```go
func PrintCurrentCLISettings()
```

#### func  PrintUserWaitTime

```go
func PrintUserWaitTime(user string)
```

#### func  ProcessExistingURL

```go
func ProcessExistingURL(longURL string)
```

#### func  SetInvalidSettingsToDefaults

```go
func SetInvalidSettingsToDefaults()
```
used regularly in CLI to constantly provide valid settings

#### func  Start

```go
func Start() *redis.Client
```
