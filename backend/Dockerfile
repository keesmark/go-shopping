FROM golang:1.16.4-alpine

COPY . /app

WORKDIR /app

RUN apk update \
  && apk add --no-cache git \
  && go get github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql

EXPOSE 8080