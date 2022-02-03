#!/usr/bin/make -f

export CGO_ENABLED=0

# Builds the project.
build:
	$(call go_build,linux,amd64)
	$(call go_build,darwin,amd64)
	$(call go_build,darwin,arm64)

# Builds the project.
define go_build
	GOOS=${1} GOARCH=${2} go build -o bin/skpr-s3-sync_${1}_${2} -ldflags='-extldflags "-static"' github.com/skpr/s3-sync
endef

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Run golint against code
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

.PHONY: *
