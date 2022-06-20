FROM golang:1.18

WORKDIR /niqurl
COPY . .
WORKDIR /niqurl/cmd/cli/
RUN go get -u && go mod tidy & go build -o niqurl
