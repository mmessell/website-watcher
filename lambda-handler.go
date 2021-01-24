package main

import (
	"context"

	"./application"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return application.Start()
}

func main() {
	lambda.Start(HandleRequest)
}
