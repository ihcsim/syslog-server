package main

import (
	"log"
	"log/syslog"
)

type logAPI func(string) error

func main() {
	host := "192.168.99.100:10154"
	client, err := syslog.Dial("tcp", host, syslog.LOG_INFO, "")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	api := []logAPI{
		client.Alert, client.Crit, client.Debug,
		client.Emerg, client.Err, client.Info,
		client.Notice, client.Warning,
	}

	msg := "Hello World!"
	for _, a := range api {
		if err := a(msg); err != nil {
			log.Print(err)
		}
	}
}
