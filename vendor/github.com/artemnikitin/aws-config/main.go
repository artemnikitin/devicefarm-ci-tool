package awsconfig

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

var (
	logging = flag.Bool("log", false, "Enable logging")
	region  = flag.String("region", "us-east-1", "Set AWS region")
	role    = flag.String("role", "", "Role ARN")
)

// New creates config with credentials for using with AWS SDK for Go
func New() *aws.Config {
	config := aws.NewConfig()
	if *role != "" {
		session := session.New(config)
		client := sts.New(session)
		resp, err := client.GetSessionToken(&sts.GetSessionTokenInput{})
		if err != nil {
			log.Fatal("Can't retrieve temprorarily credentials because of:", err)
		}
		config.WithCredentials(credentials.NewStaticCredentials(
			*resp.Credentials.AccessKeyId,
			*resp.Credentials.SecretAccessKey,
			*resp.Credentials.SessionToken))
	} else {
		config.WithCredentials(credentials.NewEnvCredentials())
	}
	config.WithRegion(*region)
	if *logging {
		config.WithLogLevel(aws.LogDebugWithHTTPBody)
	}
	return config
}
