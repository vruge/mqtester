#!/bin/bash

cp ../../certs/server/hivemq.jks ./hivemq.jks
podman build -t hivemq_tls -f Dockerfile --platform=linux/amd64
podman kill $(podman ps -q) || echo "no kill"
podman run -p 8080:8080 -p 1883:1883 -p 8883:8883 -p 8081:8081 localhost/hivemq_tls
