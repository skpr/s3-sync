FROM golang:1.23-alpine as build
ADD . /go/src/github.com/skpr/s3-sync
WORKDIR /go/src/github.com/skpr/s3-sync
RUN apk add make
RUN make build

FROM docker.io/amazon/aws-cli:2.17.59
COPY --from=build /go/src/github.com/skpr/s3-sync/bin/s3-sync /usr/local/bin/s3-sync
ENTRYPOINT ["/usr/local/bin/s3-sync"]
