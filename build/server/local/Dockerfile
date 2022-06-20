FROM golang:1.18

WORKDIR /niqurl
COPY . .
WORKDIR /niqurl/cmd/server
RUN go build -o server
EXPOSE ${SERVER_PORT}