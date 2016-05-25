#!/bin/bash

docker run --rm \
  --log-driver=syslog \
  --log-opt syslog-address=tcp://192.168.99.100:10154 \
  --log-opt syslog-format=rfc5424 \
  ubuntu:wily \
  /bin/bash -c 'for ((i = 0; ; i++)); do echo "$i: $(date)"; sleep 1; done'
