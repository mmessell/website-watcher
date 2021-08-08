package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmessell/website-watcher/business"
)

func HandleLambdaEvent() (business.MyResponse, error) {
	return business.Run()
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
