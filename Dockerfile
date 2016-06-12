FROM golang:1.6-wheezy
MAINTAINER Tristan S.

ENV APP_DIR="/usr/sudoku_web"

RUN mkdir -p $APP_DIR
WORKDIR  $APP_DIR

COPY public $APP_DIR/public
COPY src $APP_DIR/src
COPY templates $APP_DIR/templates

RUN go get github.com/gorilla/mux
RUN go get github.com/jcelliott/lumber
RUN go get github.com/tri125/sudoku

RUN go build $APP_DIR/src/main/main.go

EXPOSE 4040
CMD [ "sh", "-c", "go run ${APP_DIR}/src/main/main.go"]