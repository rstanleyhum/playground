FROM openjdk:8

LABEL maintainer="r.stanley.hum@gmail.com"

RUN apt-get update -y && apt-get install -y icedtea-8-plugin && apt-get upgrade -y

CMD ["/bin/bash"]

#
# This Docker allows for using Java and JavaWebStart (javaws) without installing
# Java
#
# docker run -it --rm -v C:/Users/stanley/Downloads:/downloads -e DISPLAY=<ip>:0 rstanleyhum/java javaws /downloads/launch.jnlp
#
# Need to run VcXsrv
#
