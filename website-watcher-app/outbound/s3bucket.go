package outbound

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"strings"
)

type Bucket struct {
	name   string
	region string
}

func NewBucket(name string, region string) Bucket {
	return Bucket{name: name, region: region}
}

func (bucket Bucket) PutObject(key string, state string) error {
	svc := bucket.initSession()

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket.name),
		Key:    aws.String(key),
		Body:   aws.ReadSeekCloser(strings.NewReader(state)),
	})

	return err
}

func (bucket Bucket) GetObject(key string) ([]byte, error) {
	svc := bucket.initSession()

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket.name),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (bucket Bucket) initSession() *s3.S3 {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(bucket.region)},
	)

	return s3.New(sess)
}
