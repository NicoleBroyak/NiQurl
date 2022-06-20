FROM golang:1.18

WORKDIR /niqurl
COPY . .
WORKDIR /niqurl/deployments/local
RUN . ./gomodupdate.sh
WORKDIR /niqurl/cmd/cli/
RUN go build -o niqurl
