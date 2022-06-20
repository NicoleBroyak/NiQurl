module github.com/nicolebroyak/niqurl/tools/redishandler

go 1.18

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/nicolebroyak/niqurl/config/niqurlconfigs v0.0.0-20220620114441-90dcc89fc627
	github.com/nicolebroyak/niqurl/tools/randomusers v0.0.0-20220620121358-59b6bccadad0
	github.com/nicolebroyak/niqurl/tools/urlhandler v0.0.0-20220620114441-90dcc89fc627
)

replace github.com/nicolebroyak/niqurl/config/niqurlconfigs v0.0.0-20220620114441-90dcc89fc627 => ../../config/niqurlconfigs

replace github.com/nicolebroyak/niqurl/tools/randomusers v0.0.0-20220620114441-90dcc89fc627 => ../randomusers

replace github.com/nicolebroyak/niqurl/tools/urlhandler v0.0.0-20220620114441-90dcc89fc627 => ../urlhandler

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)
