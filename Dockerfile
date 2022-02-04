FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY s3-sync /usr/local/bin/
RUN chmod +x /usr/local/bin/s3-sync
ENTRYPOINT ["/usr/local/bin/s3-sync"]