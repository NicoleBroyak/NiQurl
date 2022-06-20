FROM golang:1.18-alpine

WORKDIR /niqurl
COPY . .
WORKDIR /niqurl/deployments/local
RUN . ./gomodupdate.sh
WORKDIR /niqurl/cmd/server
RUN go build -o server
RUN go get -v github.com/cosmtrek/air
ENTRYPOINT air
EXPOSE ${SERVER_PORT}