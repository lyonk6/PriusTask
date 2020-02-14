FROM golang

LABEL maintainer = "kennethflyon@gmail.com"

# Set up:
COPY certs priustask/certs/
COPY app/src/model src/model/
COPY app/src/main.go priustask/
COPY params priustask/
RUN go get github.com/lib/pq
EXPOSE 3000 3006
WORKDIR priustask/

# Start: 
RUN go build -o pt main.go 
CMD ["./pt"]


