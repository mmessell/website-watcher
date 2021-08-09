package outbound

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

type Bucket struct {
	name   string
	region string
}

func NewBucket(name string, region string) Bucket {
	fmt.Println("NewBucket:name:" + name)
	fmt.Println("NewBucket:region:" + region)

	return Bucket{name: name, region: region}
}

func (bucket Bucket) GetObject(key string) ([]byte, bool) {
	svc := bucket.initSession()

	req, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket.name),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, false
	}

	if bytes, err := io.ReadAll(req.Body); err == nil {
		return bytes, true
	}

	return nil, false
}

func (bucket Bucket) initSession() *s3.S3 {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(bucket.region)},
	)

	return s3.New(sess)
}
