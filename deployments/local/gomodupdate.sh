cd /niqurl/api
go clean -modcache
go get -u
go mod tidy
cd /niqurl/cmd/cli
go clean -modcache
go get -u
go mod tidy
cd /niqurl/cmd/server
go clean -modcache
go get -u
go mod tidy
cd /niqurl/pkg/cli
go clean -modcache
go get -u
go mod tidy
cd /niqurl/tools/redishandler
go clean -modcache
go get -u
go mod tidy
cd /niqurl/tools/randomusers
go clean -modcache
go get -u
go mod tidy
cd /niqurl/tools/urlhandler
go clean -modcache
go get -u
go mod tidy
cd /niqurl/config/niqurlconfigs
go clean -modcache
go get -u
go mod tidy