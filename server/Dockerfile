FROM golang

ADD restchat-server /usr/src/restchat-server

WORKDIR /usr/src/restchat-server

RUN go build .

CMD go run .

