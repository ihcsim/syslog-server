#!/bin/bash

echo -e "\033[32m10 messages in RFC5424\033[m"
docker run --rm \
  --log-driver=syslog \
  --log-opt syslog-address=tcp://192.168.99.100:10514 \
  --log-opt syslog-format=rfc5424 \
  ubuntu:wily \
  /bin/bash -c 'for ((i = 0; i <= 10; i++)); do echo "$i: $(date)"; sleep 1; done'

echo -e "\033[32m10 messages in RFC3164\033[m"
docker run --rm \
  --log-driver=syslog \
  --log-opt syslog-address=tcp://192.168.99.100:10514 \
  --log-opt syslog-format=rfc3164 \
  ubuntu:wily \
  /bin/bash -c 'for ((i = 0; i <=10 ; i++)); do echo "$i: $(date)"; sleep 1; done'
