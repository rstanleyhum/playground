FROM microsoft/mssql-server-linux:2017-CU2

LABEL maintainer="r.stanley.hum@gmail.com"

RUN apt-get -y update && apt-get -y upgrade && apt-get -y install git

RUN mkdir /working

WORKDIR /working

RUN git clone https://github.com/OHDSI/CommonDataModel.git

COPY . /working

RUN chmod +x ./setupdb.sh

CMD /bin/bash ./entrypoint.sh
