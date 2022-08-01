# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY . ./

RUN go build cmd/server/server.go

EXPOSE 8082

CMD [ "./server" ]