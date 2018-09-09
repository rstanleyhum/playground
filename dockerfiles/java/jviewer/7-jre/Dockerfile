FROM openjdk:7

LABEL maintainer="R.Stanley Hum <r.stanley.hum@gmail.com>"

RUN apt-get update -y && apt-get install -y icedtea-7-plugin && apt-get upgrade -y

CMD ["/bin/bash"]

#
# This Docker allows for using Java and JavaWebStart (javaws) without installing
# Java
#
# docker run -it --rm -v C:/Users/stanley/Downloads:/downloads -e DISPLAY=<ip>:0 rstanleyhum/java javaws /downloads/launch.jnlp
#
# Need to run VcXsrv
#
