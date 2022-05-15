FROM golang:1.16-alpine

WORKDIR /opt/code
ADD ./ /opt/code/

RUN apc update && apc upgrade && \
apc add --no-cahe git

RUN go mod download 

RUN go build -o bin/mock-service cmd/mock-service/main.go
ENTRYPOINT ["bin/mock-service"]