FROM ubuntu:xenial

LABEL maintainer="r.stanley.hum@gmail.com"

RUN apt-get update -qq \
 && apt-get install -y bash git curl unzip lib32stdc++6 wget \
 && apt-get upgrade -y 

RUN git clone -b beta https://github.com/flutter/flutter.git \
 && export PATH=`pwd`/flutter/bin:$PATH \
 && flutter doctor

RUN wget https://dl.google.com/android/repository/sdk-tools-linux-3859397.zip


CMD ["/bin/bash"]


