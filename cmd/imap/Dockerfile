FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/shauncampbell/unified-ldap/
COPY . .

RUN go build -o /go/bin/ldap ./cmd/ldap

FROM alpine:3.12

COPY --from=builder /go/bin/ldap /go/bin/ldap
LABEL maintainer="Shaun Campbell <docker@shaun.scot>"

VOLUME /etc/ldap.yaml

ENTRYPOINT ["./go/bin/ldap", "-f", "/etc/ldap.yaml"]