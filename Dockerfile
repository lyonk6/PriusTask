FROM golang

LABEL maintainer = "kennethflyon@gmail.com"

# Set up:
COPY app/src/model src/model/
COPY app/src/main.go priustask/
COPY certs priustask/certs/
RUN go get github.com/lib/pq


## Expose these ports for https and postgresql:
EXPOSE 8080 5432 22
WORKDIR priustask/

# Build the application:
RUN go build -o pt main.go
CMD ./pt
