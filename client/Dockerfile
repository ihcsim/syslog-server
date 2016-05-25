FROM golang:1.6.2-alpine
MAINTAINER Ivan Sim, ihcsim@gmail.com

COPY . $GOPATH/src/github.com/ihcsim/go-syslog/client
WORKDIR $GOPATH/src/github.com/ihcsim/go-syslog/client
RUN go install -v .
ENTRYPOINT ["client"]
