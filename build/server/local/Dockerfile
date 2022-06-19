FROM golang:1.18

WORKDIR /go/app/
COPY . .
WORKDIR /go/app/cmd/server/
RUN go mod tidy
RUN go build -o server

EXPOSE ${REDIS_PORT}