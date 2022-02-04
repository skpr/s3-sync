FROM alpine:latest

RUN apk add --no-cache ca-certificates

RUN adduser -D -u 1000 skpr

COPY s3-sync /usr/local/bin/
RUN chmod +x /usr/local/bin/s3-sync

USER skpr

ENTRYPOINT ["/usr/local/bin/s3-sync"]
