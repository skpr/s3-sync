#!/usr/bin/make -f

export CGO_ENABLED=0

#!/usr/bin/make -f

lint:
	revive -set_exit_status -exclude=./vendor/... ./...

test:
	go test ./...

build:
	go build -a -o bin/s3-sync github.com/skpr/s3-sync

.PHONY: *