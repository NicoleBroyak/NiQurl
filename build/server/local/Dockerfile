FROM golang:1.18

WORKDIR /niqurl
COPY . .
WORKDIR /niqurl/deployments/local
RUN . ./gomodupdate.sh
WORKDIR /niqurl/cmd/server
RUN go build -o server
EXPOSE ${SERVER_PORT}