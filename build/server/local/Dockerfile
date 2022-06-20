FROM golang:1.18

WORKDIR /niqurl
COPY . .
WORKDIR /niqurl/cmd/server
RUN go get -u && go mod tidy & go build -o server
EXPOSE ${SERVER_PORT}