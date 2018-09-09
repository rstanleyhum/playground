FROM openjdk:8-jre-alpine

LABEL maintainer="r.stanley.hum@gmail.com"

ENV LC_CTYPE en_US.UTF-8

RUN apk update && apk upgrade && apk add icedtea-web ttf-dejavu

COPY ./jviewer.jnlp /usr/src/myapp/
WORKDIR /usr/src/myapp

CMD ["/bin/sh"]
#CMD ["javaws", "./jviewer.jnlp"]
