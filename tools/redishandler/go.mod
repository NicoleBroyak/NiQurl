module github.com/nicolebroyak/niqurl/tools/redishandler

go 1.18

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/nicolebroyak/niqurl/config/niqurlconfigs v0.0.0-20220619174626-21d7a352580e
	github.com/nicolebroyak/niqurl/tools/randomusers v0.0.0-20220619172856-62cf50ad86ab
	github.com/nicolebroyak/niqurl/tools/urlhandler v0.0.0-20220619172856-62cf50ad86ab
)

replace github.com/nicolebroyak/niqurl/config/niqurlconfigs v0.0.0-20220619162436-c3bacffbf7e6 => ../../config/niqurlconfigs

replace github.com/nicolebroyak/niqurl/tools/randomusers v0.0.0-20220619162436-c3bacffbf7e6 => ../randomusers

replace github.com/nicolebroyak/niqurl/tools/urlhandler v0.0.0-20220619162436-c3bacffbf7e6 => ../urlhandler

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)
