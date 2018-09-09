FROM ubuntu:16.04

RUN apt-get update && apt-get install -y openconnect iptables expect

ADD ./connect.sh /root

RUN chmod +x /root/connect.sh

CMD ["/root/connect.sh"]
