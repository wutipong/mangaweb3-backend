#Stage 1 -- building executable
FROM golang:1.21-alpine AS builder1

WORKDIR /go/src/mangaweb
COPY . .

ARG VERSION=Development
RUN apk add git
RUN go get -d -v ./...
RUN go build -v -ldflags="-X 'main.versionString=$VERSION' " -o mangaweb .

# Stage 3 -- combine the two
FROM alpine:latest

WORKDIR /root/
COPY --from=builder1 /go/src/mangaweb/mangaweb ./
COPY --from=builder1 /go/src/mangaweb/docs/swagger.json ./docs/swagger.json
COPY --from=builder1 /go/src/mangaweb/docs/swagger.yaml ./docs/swagger.yaml

CMD ["./mangaweb"]