package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmessell/website-watcher/business"
	"github.com/mmessell/website-watcher/outbound"
	"os"
)

func HandleLambdaEvent() (string, error) {
	bucket := os.Getenv("BUCKET")
	region := os.Getenv("REGION")

	repo := outbound.NewWebsiteRepoImpl(bucket, region, "users-and-websites.json")
	hc := outbound.HttpClientImpl{}
	ww := business.NewWebsiteWatcher(repo, hc)
	err := ww.Run()

	if err != nil {
		return "ERROR", err
	} else {
		return "SUCCESS", err
	}
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
