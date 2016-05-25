FROM golang:1.6.2-alpine
MAINTAINER Ivan Sim, ihcsim@gmail.com

COPY . $GOPATH/src/github.com/ihcsim/go-syslog
WORKDIR $GOPATH/src/github.com/ihcsim/go-syslog
RUN go install -v .
EXPOSE 10154
ENTRYPOINT ["go-syslog"]
