package redishandler

var client = Start("niqurl-redis:6379")
var context = client.Context()
