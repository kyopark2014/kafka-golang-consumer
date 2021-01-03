###################
##  build stage  ##
###################
FROM golang:1.13.0-alpine as builder
WORKDIR /kafka-golang-consumer
COPY . .
RUN go build -v -o kafka-golang-consumer

##################
##  exec stage  ##
##################
FROM alpine:3.10.2
WORKDIR /app
COPY ./configs/config.json.default ./configs/config.json
COPY --from=builder /kafka-golang-consumer /app/
CMD ["./kafka-golang-consumer"]
