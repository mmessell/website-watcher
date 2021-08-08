package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmessell/website-watcher/business"
	"github.com/mmessell/website-watcher/outbound"
)

func HandleLambdaEvent() (business.MyResponse, error) {
	ww := business.NewWebsiteWatcher(outbound.WatchConfigRepoImpl{})
	return ww.Run()
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
