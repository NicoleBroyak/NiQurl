FROM golang:1.18

WORKDIR /go/app
RUN git clone "https://github.com/nicolebroyak/niqurl/"
WORKDIR /go/app/niqurl/cmd/cli
RUN go mod tidy
RUN go build -o niqurl
RUN source ../../config/niqurlconfigs/envfile.env

EXPOSE 8081
