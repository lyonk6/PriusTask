FROM golang

LABEL maintainer = "kennethflyon@gmail.com"

COPY certs priustask/certs/
COPY app/src/model src/model/
COPY app/src/main.go priustask/
COPY params  priustask/

RUN go get github.com/lib/pq

WORKDIR priustask/

EXPOSE 3000 3006

