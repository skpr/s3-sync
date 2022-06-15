package main

import (
	"os"
	"os/exec"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cliRegion   = kingpin.Flag("region", "When transferring objects from an s3 bucket to an s3 bucket, this specifies the region of the source bucket").Default("ap-southeast-2").String()
	cliEndpoint = kingpin.Flag("endpoint", "Override command's default URL with the given URL").Envar("SKPR_S3_SYNC_ENDPOINT").String()
	cliDelete   = kingpin.Flag("delete", "Delete files which are not listed in the source").Envar("SKPR_S3_SYNC_DELETE").Bool()
	cliExclude  = kingpin.Flag("exclude", "Exclude paths from the list to be synced").Envar("SKPR_S3_SYNC_EXCLUDE").Default(".htaccess").String()
	cliSource   = kingpin.Arg("source", "Source files which are synced (local or S3 path)").Required().String()
	cliTarget   = kingpin.Arg("target", "Target files which are synced (local or S3 path)").Required().String()
)

func main() {
	kingpin.Parse()

	args := buildArgs(*cliEndpoint, *cliSource, *cliTarget, *cliExclude)

	cmd := exec.Command("aws", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// Command which is compatible with the AWS S3 sync command line interface.
func buildArgs(endpoint, source, target, exclude string) []string {
	args := []string{"s3"}

	if endpoint != "" {
		args = append(args, "--endpoint-url", endpoint)
	}

	args = append(args, "sync")

	if exclude != "" {
		for _, e := range strings.Split(exclude, ",") {
			args = append(args, "--exclude", e)
		}
	}

	args = append(args, source, target)

	return args
}
