package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildArgsWithEndpoint(t *testing.T) {
	command := buildArgs("127.0.0.1", "foo", "bar", "")
	assert.Equal(t, []string{"s3", "--endpoint-url", "127.0.0.1", "sync", "foo", "bar"}, command)
}

func TestBuildArgsdWithExcludes(t *testing.T) {
	command := buildArgs("", "foo", "bar", "/stuff,/things")
	assert.Equal(t, []string{"s3", "sync", "--exclude", "/stuff", "--exclude", "/things", "foo", "bar"}, command)
}

func TestBuildArgsWithAll(t *testing.T) {
	command := buildArgs("127.0.0.1", "foo", "bar", "/stuff,/things")
	assert.Equal(t, []string{"s3", "--endpoint-url", "127.0.0.1", "sync", "--exclude", "/stuff", "--exclude", "/things", "foo", "bar"}, command)
}
