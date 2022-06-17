package redishandler

var client = Start()
var context = client.Context()
var serverPath string = "localhost:8081"
var settingsMap = map[string]int{
	"SHORT_URL_LEN":  4,
	"USER_WAIT_TIME": 30000,
	"URL_COUNT":      0,
	"USER_COUNT":     0,
}
