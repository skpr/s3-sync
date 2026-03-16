FROM golang:1.26-alpine AS build
ADD . /go/src/github.com/skpr/s3-sync
WORKDIR /go/src/github.com/skpr/s3-sync
RUN apk add make
RUN mise build

FROM docker.io/amazon/aws-cli:2.34.9
ARG TARGETPLATFORM
COPY $TARGETPLATFORM/skpr-s3-sync /usr/local/bin/s3-sync
ENTRYPOINT ["/usr/local/bin/s3-sync"]
