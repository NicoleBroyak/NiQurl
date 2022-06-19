FROM schadokar/redis-go-server
WORKDIR /go/app/
COPY . .
WORKDIR /go/app/cmd/cli/
EXPOSE "${REDIS_PORT}"
CMD [ "redis-server" "--port ${REDIS_PORT}" ]