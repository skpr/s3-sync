package main

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/seqsense/s3sync"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cliRegion   = kingpin.Flag("region", "When transferring objects from an s3 bucket to an s3 bucket, this specifies the region of the source bucket").Default(endpoints.ApSoutheast2RegionID).String()
	cliEndpoint = kingpin.Flag("endpoint", "Override command's default URL with the given URL").String()
	cliParallel = kingpin.Flag("parallel", "Sets maximum number of parallel file sync jobs").Default("16").Int()
	cliDelete   = kingpin.Flag("delete", "Delete files which are not listed in the source").Bool()
	cliExclude  = kingpin.Flag("exclude", "Exclude paths from the list to be synced").Envar("SKPR_S3_SYNC_EXCLUDE").String()
	cliSource   = kingpin.Arg("source", "Source files which are synced (local or S3 path)").Required().String()
	cliTarget   = kingpin.Arg("target", "Target files which are synced (local or S3 path)").Required().String()
)

func main() {
	kingpin.Parse()

	config := &aws.Config{
		Region:   cliRegion,
		Endpoint: cliEndpoint,
	}

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	var options []s3sync.Option

	if *cliDelete {
		options = append(options, s3sync.WithDelete())
	}

	if *cliParallel > 0 {
		options = append(options, s3sync.WithParallel(*cliParallel))
	}

	for _, exclude := range strings.Split(*cliExclude, ",") {
		options = append(options, s3sync.WithExcludePattern(exclude))
	}

	syncManager := s3sync.New(sess, options...)

	err = syncManager.Sync(*cliSource, *cliTarget)
	if err != nil {
		panic(err)
	}
}
