from hivemq/hivemq4

EXPOSE 1883 8000 8883 8080
COPY tls_conf.xml ./conf/config.xml
COPY hivemq.jks ./hivemq.jks
RUN cat ./conf/config.xml

# RUN rm -rf extensions/hivemq-allow-all-extension
# RUN rm -f extensions/hivemq-enterprise-security-extension/DISABLED