FROM golang

LABEL maintainer = "kennethflyon@gmail.com"

# Set up:
# COPY certs/* priustask/certs/
COPY app/src/model src/model/
COPY app/src/main.go priustask/
COPY params priustask/
RUN go get github.com/lib/pq
RUN mkdir priustask/certs

# Create our RSA key-value pair:
RUN openssl req -newkey rsa:2048 -new -nodes -x509 -days 365 -keyout priustask/certs/key.pem -out priustask/certs/cert.pem -subj '/C=US/ST=Nevada/L=Vegas/O=priustask.com/CN=PriusTask'

## Expose these ports for https and postgresql:
EXPOSE 8080 5432
WORKDIR priustask/

# Start: 
# RUN go build -o pt main.go 
# CMD ["./pt"]
