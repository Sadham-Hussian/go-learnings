FROM golang:1.15-alpine as build

RUN apk add --no-cache git

WORKDIR /src

ENV CGO_ENABLED=0 
ENV GO111MODULE=on

RUN go mod init rabbitmq-hello-world

COPY publisher.go /src

RUN go build publisher.go

FROM alpine as runtime

COPY --from=build /src/publisher /app/publisher

CMD ["/app/publisher"]