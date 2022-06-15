FROM docker.io/amazon/aws-cli:2.1.27

COPY s3-sync /usr/local/bin/
RUN chmod +x /usr/local/bin/s3-sync

RUN adduser -D -u 1000 skpr
USER skpr

ENTRYPOINT ["/usr/local/bin/s3-sync"]
