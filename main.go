package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/seqsense/s3sync"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cliRegion   = kingpin.Flag("region", "When transferring objects from an s3 bucket to an s3 bucket, this specifies the region of the source bucket").Default(endpoints.ApSoutheast2RegionID).String()
	cliEndpoint = kingpin.Flag("endpoint", "Override command's default URL with the given URL").String()
	cliParallel = kingpin.Flag("parallel", "Sets maximum number of parallel file sync jobs.").Default("4").Int()
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

	syncManager := s3sync.New(sess, s3sync.WithParallel(*cliParallel), s3sync.WithDelete())

	err = syncManager.Sync(*cliSource, *cliTarget)
	if err != nil {
		panic(err)
	}
}
