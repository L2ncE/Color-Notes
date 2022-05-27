FROM golang:latest

MAINTAINER YXH

RUN mkdir -p /data/note
WORKDIR /data/note
COPY . /data/note

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
RUN go mod download

RUN go build main.go

EXPOSE 5556

//ENTRYPOINT  ["./main"]