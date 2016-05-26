package main

import (
	"log"
	"log/syslog"
	"os"
)

type logAPI func(string) error

func main() {
	host := "127.0.0.1:10514"
	if os.Getenv("SERVER_HOST") != "" {
		host = os.Getenv("SERVER_HOST")
	}

	log.Print("Sending messages to ", host)
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
