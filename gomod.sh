cd /home/nicolebroyak/code/src/github.com/nicolebroyak/niqurl/api
go clean -modcache
go get -u
go mod tidy
cd /home/nicolebroyak/code/src/github.com/nicolebroyak/niqurl/cmd/cli
go clean -modcache
go get -u
go mod tidy
cd /home/nicolebroyak/code/src/github.com/nicolebroyak/niqurl/cmd/server
go clean -modcache
go get -u
go mod tidy
cd /home/nicolebroyak/code/src/github.com/nicolebroyak/niqurl/pkg/cli
go clean -modcache
go get -u
go mod tidy
cd /home/nicolebroyak/code/src/github.com/nicolebroyak/niqurl/tools/redishandler
go clean -modcache
go get -u
go mod tidy
cd /home/nicolebroyak/code/src/github.com/nicolebroyak/niqurl/tools/randomusers
go clean -modcache
go get -u
go mod tidy
cd /home/nicolebroyak/code/src/github.com/nicolebroyak/niqurl/tools/urlhandler
go clean -modcache
go get -u
go mod tidy
cd /home/nicolebroyak/code/src/github.com/nicolebroyak/niqurl/config/niqurlconfigs
go clean -modcache
go get -u
go mod tidy
cd /home/nicolebroyak/code/src/github.com/nicolebroyak/niqurl/
