// Package main wraps the AWS CLI command line utility for syncing files.
package main

import (
	"log/slog"
	"os"
	"os/exec"
	"strings"

	"github.com/alecthomas/kingpin/v2"
)

var (
	cliEndpoint = kingpin.Flag("endpoint", "Override command's default URL with the given URL").Envar("SKPR_S3_SYNC_ENDPOINT").String()
	cliExclude  = kingpin.Flag("exclude", "Exclude paths from the list to be synced").Envar("SKPR_S3_SYNC_EXCLUDE").Default(".htaccess").String()
	cliMode     = kingpin.Flag("mode", "Mode which will used for syncing (sync or s3)").Envar("SKPR_S3_SYNC_MODE").Default("sync").String()
	cliSource   = kingpin.Arg("source", "Source files which are synced (local or S3 path)").Required().String()
	cliTarget   = kingpin.Arg("target", "Target files which are synced (local or S3 path)").Required().String()
)

func main() {
	kingpin.Parse()

	args := buildArgs(*cliEndpoint, *cliMode, *cliSource, *cliTarget, *cliExclude)

	slog.Info("Starting sync", "args", strings.Join(args, " "))

	cmd := exec.Command("aws", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("Sync finished")
}

// Command which is compatible with the AWS S3 sync command line interface.
func buildArgs(endpoint, mode, source, target, exclude string) []string {
	args := []string{"s3"}

	if endpoint != "" {
		args = append(args, "--endpoint-url", endpoint)
	}

	args = append(args, mode)

	if exclude != "" {
		for _, e := range strings.Split(exclude, ",") {
			args = append(args, "--exclude", e)
		}
	}

	args = append(args, source, target)

	return args
}
