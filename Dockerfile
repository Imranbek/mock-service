FROM golang:1.16-alpine as builder

WORKDIR /opt/code
COPY . .

RUN apk update && apk upgrade && apk add --no-cache git

RUN go mod download 

RUN go build -o /mock-service cmd/mock-service/main.go

FROM alpine:3.15 as app

COPY --from=builder /mock-service /usr/local/bin

ENTRYPOINT ["mock-service"]
