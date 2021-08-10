package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmessell/website-watcher/business"
	"github.com/mmessell/website-watcher/outbound"
	"os"
)

func HandleLambdaEvent() (bool, error) {
	bucket := os.Getenv("BUCKET")
	region := os.Getenv("REGION")

	repo := outbound.NewWebsiteRepoImpl(bucket, region, "users-and-websites.json")
	ww := business.NewWebsiteWatcher(repo)
	return ww.Run(), nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
