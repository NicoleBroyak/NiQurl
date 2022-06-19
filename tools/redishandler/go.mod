module github.com/nicolebroyak/niqurl/tools/redishandler

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/nicolebroyak/niqurl/config/niqurlconfigs v0.0.0
	github.com/nicolebroyak/niqurl/tools/randomusers v0.0.0
	github.com/nicolebroyak/niqurl/tools/urlhandler v0.0.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)

replace github.com/nicolebroyak/niqurl/config/niqurlconfigs v0.0.0 => ../../config/niqurlconfigs

replace github.com/nicolebroyak/niqurl/tools/randomusers v0.0.0 => ../randomusers

replace github.com/nicolebroyak/niqurl/tools/urlhandler v0.0.0 => ../urlhandler

go 1.18
