# syntax=docker/dockerfile:1
# IT IS NOT FINISHED YET
FROM golang:alpine as dev

RUN apk update && apk add curl
WORKDIR /usr/local/share/ca-certificates
RUN apk add ca-certificates
RUN /usr/sbin/update-ca-certificates

WORKDIR /app/

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api ./src/main.go

EXPOSE 8080
EXPOSE 8081

CMD ["./api"]