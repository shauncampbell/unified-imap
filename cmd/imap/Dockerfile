FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/shauncampbell/unified-imap/
COPY . .

RUN go build -o /go/bin/imap ./cmd/imap

FROM alpine:3.12

COPY --from=builder /go/bin/imap /go/bin/imap
LABEL maintainer="Shaun Campbell <docker@shaun.scot>"

VOLUME /etc/imap.yaml

ENTRYPOINT ["./go/bin/imap", "-f", "/etc/imap.yaml"]