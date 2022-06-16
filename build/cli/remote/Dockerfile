FROM golang:1.18

WORKDIR /go/app
RUN git clone "https://github.com/nicolebroyak/niqurl/"
WORKDIR /go/app/niqurl/cmd/cli
RUN go mod tidy
RUN go build -o niqurl

EXPOSE 8081
CMD [ "./niqurl" ]