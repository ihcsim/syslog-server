FROM golang:1.6.2-alpine
MAINTAINER Ivan Sim, ihcsim@gmail.com

ARG PACKAGE=github.com/ihcsim/go-syslog
ENV HOST ":10514"

COPY . $GOPATH/src/$PACKAGE
WORKDIR $GOPATH/src/$PACKAGE
RUN go install -v .
EXPOSE 10154
ENTRYPOINT ["go-syslog"]
