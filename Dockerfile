# build stage
FROM golang:alpine AS builder

RUN apk add git gcc libc-dev

RUN mkdir /go/src/coupon-service 
WORKDIR /go/src/coupon-service

COPY . .

RUN go build ./cmd/coupon_service/main.go

ENTRYPOINT ./main
EXPOSE 80
