package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildArgsWithEndpoint(t *testing.T) {
	command, err := buildArgs("127.0.0.1", "sync", "foo", "bar", "")
	assert.NoError(t, err)
	assert.Equal(t, []string{"s3", "--endpoint-url", "127.0.0.1", "sync", "foo", "bar"}, command)
}

func TestBuildArgsWithEndpointAndCP(t *testing.T) {
	command, err := buildArgs("127.0.0.1", "cp", "foo", "bar", "")
	assert.NoError(t, err)
	assert.Equal(t, []string{"s3", "--endpoint-url", "127.0.0.1", "cp", "--recursive", "foo", "bar"}, command)
}

func TestBuildArgsdWithExcludes(t *testing.T) {
	command, err := buildArgs("", "sync", "foo", "bar", "/stuff,/things")
	assert.NoError(t, err)
	assert.Equal(t, []string{"s3", "sync", "--exclude", "/stuff", "--exclude", "/things", "foo", "bar"}, command)
}

func TestBuildArgsWithAll(t *testing.T) {
	command, err := buildArgs("127.0.0.1", "sync", "foo", "bar", "/stuff,/things")
	assert.NoError(t, err)
	assert.Equal(t, []string{"s3", "--endpoint-url", "127.0.0.1", "sync", "--exclude", "/stuff", "--exclude", "/things", "foo", "bar"}, command)
}
