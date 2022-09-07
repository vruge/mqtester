```shell
keytool -genkey -keyalg RSA -alias hivemq -keystore hivemq.jks -storepass betterChangeMe -validity 360 -keysize 2048 -dname "CN=localhost, OU=Foo, O=Bar, L=City, ST=State, C=DE"

keytool -exportcert -keystore hivemq.jks -alias hivemq -keypass betterChangeMe -storepass betterChangeMe -rfc -file hivemq-server-cert.pem 
keytool -export -keystore hivemq.jks -alias hivemq -storepass betterChangeMe -file hivemq-server.crt

```


```shell
keytool -import -file hivemq-server.crt -alias HiveMQ -keystore mqtt-client-trust-store.jks -storepass betterChangeMe
keytool -exportcert -alias hivemq -keypass betterChangeMe -storepass betterChangeMe -keystore mqtt-client-trust-store.jks -rfc -file mqtt-client-trust-store.pem
```