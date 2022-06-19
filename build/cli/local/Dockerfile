FROM golang:1.18

WORKDIR /go/app/
COPY . .

WORKDIR /go/app/cmd/cli/
RUN . ../../config/niqurlconfigs/envfile.sh
RUN go mod tidy
RUN go build -o niqurl

EXPOSE 8081
