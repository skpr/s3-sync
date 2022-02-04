FROM alpine:latest

RUN apk add --no-cache ca-certificates bash

RUN adduser -D -u 1000 skpr

COPY s3-sync /usr/local/bin/
RUN chmod +x /usr/local/bin/s3-sync

COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

USER skpr

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
