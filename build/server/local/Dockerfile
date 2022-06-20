FROM golang:1.18

WORKDIR /niqurl
COPY . .
WORKDIR /
RUN cp /niqurl/config/niqurlconfigs/init.sh init.sh
RUN chmod +x init.sh
RUN . /init.sh
EXPOSE ${SERVER_PORT}