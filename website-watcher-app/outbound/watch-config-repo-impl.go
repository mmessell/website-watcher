package outbound

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mmessell/website-watcher/business"
	"io"
)

type WatchConfigRepoImpl struct {
}

func (r WatchConfigRepoImpl) ListAll() []business.Person {
	bucket := "website-watcher-bucket"

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	req, _ := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("users-and-websites.json"),
	})

	var persons []business.Person
	if b, err := io.ReadAll(req.Body); err == nil {
		json.Unmarshal(b, &persons)
	}

	return persons
}
