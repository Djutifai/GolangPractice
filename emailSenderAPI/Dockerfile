# syntax=docker/dockerfile:1
# IT IS NOT FINISHED YET
FROM golang:1.17-alpine

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o emailsender ./src/main.go

EXPOSE 8080
EXPOSE 8081

CMD ["/emailsender"]