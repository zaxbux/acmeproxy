FROM golang:1-alpine as builder

RUN apk --update upgrade \
&& apk --no-cache --no-progress add make git \
&& rm -rf /var/cache/apk/*

WORKDIR /go/src/github.com/zaxbux/acmeproxy
COPY . .
RUN make build

FROM alpine:3.12
RUN apk update && apk add --no-cache --virtual ca-certificates
COPY --from=builder /go/src/github.com/zaxbux/acmeproxy/dist/acmeproxy /usr/bin/acmeproxy
ENTRYPOINT [ "/usr/bin/acmeproxy" ]
