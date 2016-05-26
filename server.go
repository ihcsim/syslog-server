package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mcuadros/go-syslog.v2"
	"gopkg.in/mcuadros/go-syslog.v2/format"
)

var server *syslog.Server

func main() {
	server = syslog.NewServer()
	defer killServer()

	f := &format.Automatic{}
	server.SetFormat(f)

	logChan := make(syslog.LogPartsChannel)
	go recvLogs(logChan)
	setServerLogHandler(logChan)

	errChan := make(chan error)
	go recvErrors(errChan)

	host := "127.0.0.1:10514"
	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	}
	if err := server.ListenTCP(host); err != nil {
		log.Fatal(err)
	}

	if err := server.Boot(); err != nil {
		log.Fatal(err)
	}

	log.Print("Starting server at ", host)
	server.Wait()
}

func killServer() {
	if err := server.Kill(); err != nil {
		log.Print(err)
	}
}

func recvLogs(logChan syslog.LogPartsChannel) {
	for logPart := range logChan {
		b := &bytes.Buffer{}
		for k, v := range logPart {
			var value string
			switch v := v.(type) {
			case string:
				value = v
			case time.Time:
				value = v.String()
			case int:
				value = strconv.Itoa(v)
			default:
				log.Printf("Unexpected type %T", v)
				break
			}

			if value == "" {
				value = `""`
			}

			log := strings.Title(k) + ": " + value + "|"
			_, err := b.WriteString(log)
			if err != nil {

			}
		}

		s := b.String()
		log.Print(s[:len(s)-1])
		log.Print(logPart)
	}
}

func setServerLogHandler(logChan syslog.LogPartsChannel) {
	handler := syslog.NewChannelHandler(logChan)
	server.SetHandler(handler)
}

func recvErrors(errChan chan error) {
	for err := range errChan {
		log.Print(err)
	}
}
