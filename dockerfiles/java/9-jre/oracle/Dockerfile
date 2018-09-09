FROM alpine:3.7

LABEL maintainer="R. Stanley Hum <r.stanley.hum@gmail.com>"

# Install dependencies
RUN apk --update add curl ca-certificates tar

# Java version
ENV JAVA_VERSION_MAJOR 9
ENV JAVA_VERSION_MINOR 0
ENV JAVA_VERSION_BUILD 4
ENV JAVA_PACKAGE       jre

RUN mkdir /opt


CMD ["/bin/sh"]
