FROM golang:1.10.2-alpine3.7 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/s-take/goecho-postgre-sample

COPY Gopkg.lock Gopkg.toml ./
COPY vendor vendor
COPY db db
COPY schema schema
COPY halder handler
COPY main.go main.go

RUN go install ./...

FROM alpine:3.7
WORKDIR /usr/bin
COPY --from=build /go/bin .
