# go-syslog

[![Codeship Status for ihcsim/syslog-server](https://app.codeship.com/projects/a0e76fb0-b665-0134-c236-7ea8c0f9c13a/status?branch=master)](https://app.codeship.com/projects/194350)

This is a sample project where a client is used to send syslog messages to a syslog server. The objective is to explore the Go standard [syslog](https://golang.org/pkg/log/syslog/) library and the Docker [syslog log driver](https://docs.docker.com/engine/admin/logging/overview/).

The following is a list of syslog API that the client uses to send different message types to the server:

1. [syslog.Alert](https://golang.org/pkg/log/syslog/#Writer.Alert)
1. [syslog.Crit](https://golang.org/pkg/log/syslog/#Writer.Crit)
1. [syslog.Debug](https://golang.org/pkg/log/syslog/#Writer.Debug)
1. [syslog.Emerg](https://golang.org/pkg/log/syslog/#Writer.Alert)
1. [syslog.Err](https://golang.org/pkg/log/syslog/#Writer.Err)
1. [syslog.Info](https://golang.org/pkg/log/syslog/#Writer.Info)
1. [syslog.Notice](https://golang.org/pkg/log/syslog/#Writer.Notice)
1. [syslog.Warning](https://golang.org/pkg/log/syslog/#Writer.Warning)

## Getting Started
To build and run the syslog server docker container:

```sh
$ docker build --rm -t isim/go-syslog .
$ docker run --rm -p 10154:10154 isim/go-syslog
```

To build and run the client Docker container:

```sh
$ docker build --rm -t isim/gosyslog-client .
$ docker run --rm isim/gosyslog-client
```

The log emitter script can be used to generate a continuous stream of logs that is sent to the syslog server using the Docker syslog log driver.

```sh
$ client/emit_logs.sh
```


## LICENSE
Refer the [LICENSE](LICENSE) file.
