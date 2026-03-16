FROM docker.io/amazon/aws-cli:2.34.9
ARG TARGETPLATFORM
COPY $TARGETPLATFORM/skpr-s3-sync /usr/local/bin/s3-sync
ENTRYPOINT ["/usr/local/bin/s3-sync"]
