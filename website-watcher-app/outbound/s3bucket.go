package outbound

import (
	"fmt"
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

func (bucket Bucket) PutObject(key string, state string) {
	svc := bucket.initSession()

	//func (c *S3) PutObject(input *PutObjectInput) (*PutObjectOutput, error)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket.name),
		Key:    aws.String(key),
		Body:   aws.ReadSeekCloser(strings.NewReader(state)),
	})

	fmt.Println(err)
}

func (bucket Bucket) GetObject(key string) ([]byte, bool) {
	svc := bucket.initSession()

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket.name),
		Key:    aws.String(key),
	})

	fmt.Println(resp)

	if err != nil {
		return nil, false
	}

	s3objectBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, false
	}

	return s3objectBytes, true
}

func (bucket Bucket) initSession() *s3.S3 {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(bucket.region)},
	)

	return s3.New(sess)
}
