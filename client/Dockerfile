FROM golang:1.6.2-alpine
MAINTAINER Ivan Sim, ihcsim@gmail.com

ARG PACKAGE=github.com/ihcsim/go-syslog/client
ENV SERVER_HOST "192.168.99.100:10514"

COPY . $GOPATH/src/$PACKAGE
WORKDIR $GOPATH/src/$PACKAGE
RUN go install -v .
ENTRYPOINT ["client"]
