FROM docker.io/amazon/aws-cli:2.1.27

COPY s3-sync /usr/local/bin/
RUN chmod +x /usr/local/bin/s3-sync

ENTRYPOINT ["/usr/local/bin/s3-sync"]
