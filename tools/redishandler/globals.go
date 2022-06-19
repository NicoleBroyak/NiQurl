package redishandler

import "github.com/nicolebroyak/niqurl/config/niqurlconfigs"

var client = Start(niqurlconfigs.RedisHost)
var context = client.Context()
