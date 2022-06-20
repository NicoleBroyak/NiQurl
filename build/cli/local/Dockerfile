FROM golang:1.18-alpine

WORKDIR /niqurl
COPY . .
WORKDIR /niqurl/deployments/local
RUN . ./gomodupdate.sh
WORKDIR /niqurl/cmd/cli/
RUN go build -o niqurl
RUN go get -v github.com/cosmtrek/air
ENTRYPOINT ["air"]
