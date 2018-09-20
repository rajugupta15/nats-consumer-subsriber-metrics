FROM golang:1.10.3-alpine3.7

COPY . /go/src/github.com/rajugupta15/nats-consumer-subsriber-metrics
WORKDIR /go/src/github.com/rajugupta15/nats-consumer-subsriber-metrics
RUN apk add git

RUN go get ; CGO_ENABLED=0 GOOS=linux go build -o nats-consumer-subsriber-metrics main.go
FROM alpine:3.8
RUN apk update; apk add ca-certificates
COPY --from=0 /go/src/github.com/rajugupta15/nats-consumer-subsriber-metrics/nats-consumer-subsriber-metrics /nats-consumer-subsriber-metrics
CMD [ "/nats-consumer-subsriber-metrics" ]