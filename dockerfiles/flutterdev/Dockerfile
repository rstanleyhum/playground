FROM ubuntu:16.04

LABEL maintainer="R. Stanley Hum <r.stanley.hum@gmail.com>"

RUN apt-get -y update && apt-get -y upgrade && apt-get install -y git unzip curl lib32stdc++6 wget default-jdk

RUN mkdir /develop

WORKDIR /develop

RUN git clone -b beta https://github.com/flutter/flutter.git

RUN wget https://dl.google.com/android/repository/sdk-tools-linux-3859397.zip

RUN unzip sdk-tools-linux-3859397.zip

RUN yes | ./tools/bin/sdkmanager --licenses

RUN ./tools/bin/sdkmanager "platform-tools" "platforms;android-27" "build-tools;27.0.3" "ndk-bundle"

RUN ./tools/bin/sdkmanager --update

ENV ANDROID_HOME "/develop"
ENV JAVA_HOME "/usr/lib/jvm/java-8-openjdk-amd64/jre"
ENV PATH "/develop/flutter/bin:$PATH"

RUN flutter doctor

CMD ["/bin/bash"]

