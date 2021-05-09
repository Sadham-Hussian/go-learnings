FROM golang:1.15-alpine AS build

RUN apk add --no-cache git

WORKDIR /src

RUN go get github.com/streadway/amqp

COPY consumer.go /src

RUN go build consumer.go

FROM alpine as runtime

COPY --from=build /src/consumer /app/consumer

CMD ["/app/consumer"]