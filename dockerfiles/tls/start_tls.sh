#!/bin/bash

rm -f hivemq.jks || echo "no clean"
keytool -genkey -keyalg RSA -alias hivemq -keystore hivemq.jks -storepass betterChangeMe -validity 360 -keysize 2048 -dname "CN=localhost, OU=Foo, O=Bar, L=City, ST=State, C=DE"
podman build -t hivemq_tls -f Dockerfile --platform=linux/amd64
podman kill $(podman ps -q) || echo "no kill"
podman run -p 8080:8080 -p 1883:1883 -p 8883:8883 -p 8081:8081 localhost/hivemq_tls
