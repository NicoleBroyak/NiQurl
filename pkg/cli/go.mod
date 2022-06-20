module github.com/nicolebroyak/niqurl/pkg/cli

go 1.18

require (
	github.com/desertbit/grumble v1.1.3
	github.com/fatih/color v1.13.0
	github.com/nicolebroyak/niqurl/config/niqurlconfigs v0.0.0-20220620114441-90dcc89fc627
	github.com/nicolebroyak/niqurl/tools/randomusers v0.0.0-20220620114441-90dcc89fc627
	github.com/nicolebroyak/niqurl/tools/redishandler v0.0.0-20220620114441-90dcc89fc627
	github.com/nicolebroyak/niqurl/tools/urlhandler v0.0.0-20220620114441-90dcc89fc627
)

replace github.com/nicolebroyak/niqurl/tools/redishandler v0.0.0-20220620114441-90dcc89fc627 => ../../tools/redishandler

replace github.com/nicolebroyak/niqurl/tools/randomusers v0.0.0-20220620114441-90dcc89fc627 => ../../tools/randomusers

replace github.com/nicolebroyak/niqurl/tools/urlhandler v0.0.0-20220620114441-90dcc89fc627 => ../../tools/urlhandler

replace github.com/nicolebroyak/niqurl/config/niqurlconfigs v0.0.0-20220620114441-90dcc89fc627 => ../../config/niqurlconfigs

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/desertbit/closer/v3 v3.1.3 // indirect
	github.com/desertbit/columnize v2.1.0+incompatible // indirect
	github.com/desertbit/go-shlex v0.1.1 // indirect
	github.com/desertbit/readline v1.5.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.0.0-20220615213510-4f61da869c0c // indirect
)
