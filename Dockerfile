FROM golang:1.12.5-alpine3.9 AS base
RUN apk add --no-cache curl git

WORKDIR /go/src/github.com/VoIPGRID/dave
COPY . .
RUN go get -v

FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=base /go/bin/dave /usr/local/bin/dave
ENTRYPOINT ["/usr/local/bin/dave"]
