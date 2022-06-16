FROM golang:1.18

WORKDIR /go/app/
COPY . .
WORKDIR /go/app/cmd/server/
RUN go build -o server

EXPOSE 8081
CMD [ "./server"]