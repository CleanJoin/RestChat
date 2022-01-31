FROM golang

ADD restchat-server /usr/src/restchat-server

WORKDIR /usr/src/restchat-server

RUN go get github.com/codegangsta/gin

CMD gin . 

