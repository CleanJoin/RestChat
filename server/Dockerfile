FROM golang

ADD . /usr/src/restchat-server

WORKDIR /usr/src/restchat-server

RUN go get github.com/codegangsta/gin

RUN go build

CMD bash -c "cd /tmp && gin --path /usr/src/restchat-server"

