#   Unifi-controller
#
# 2017-12-14:
#   modified unifi-controller dockerfile directly from github.com/jessfraz/dockerfiles/unifi/Dockerfile
#   it's much better
#   reference: https://blog.jessfraz.com/post/home-lab-is-the-dopest-lab/
#
# Build:
#   docker build -t rstanleyhum/unifi .
#
# Usage:
#   docker run -d --restart always \
#       -v /etc/localtime:/etc/localtime:ro \ (probably should omit)
#       --net=host \
#       --name unifi \
#       --volume path/to/where/you/want/your/data:/config \
#       -p 3478:3478/udp \
#       -p 10001:10001/udp \
#       -p 8080:8080 \
#       -p 8081:8081 \
#       -p 8443:8443 \
#       -p 8843:8843 \
#       -p 8880:8880 \
#       rstanleyhum/unifi
#
# Remember to add an external network adapter to the HyperV moby
# also add apt-get install net-tools to find ip address
# start with: service unifi start
# https://ip address:8443/
#
# To adopt an access point and get it to show up in the software
#   ssh ubnt@$AP-IP mca-cli set-inform http://$address:8080/inform
#
FROM ubuntu:16.04

RUN echo 'deb http://www.ubnt.com/downloads/unifi/debian oldstable ubiquiti' > /etc/apt/sources.list.d/100-ubnt.list

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv 06E85760C0A52C50

RUN apt-get update && apt-get install -y unifi net-tools && apt-get -y upgrade

RUN service unifi start