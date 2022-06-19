FROM golang:1.18

WORKDIR /go/app/
COPY . .
WORKDIR /go/app/cmd/server/
RUN . ../../config/niqurlconfigs/envfile.sh
RUN go mod tidy
RUN go build -o server

EXPOSE 8081