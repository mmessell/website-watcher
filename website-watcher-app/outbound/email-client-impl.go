package outbound

import (
	"log"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	Sender  = "mmessell@me.com"
	CharSet = "UTF-8"
)

type EmailClientImpl struct {
	Region string
}

func (ec EmailClientImpl) Send(recipient string, subject string, message string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(ec.Region)},
	)

	svc := ses.New(sess)

	// Attempt to send the email.
	result, err := svc.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(message),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	})

	// Display error messages if they occur.
	if err != nil {
		log.Fatal("Couldn't send email to " + recipient)
		log.Fatal(err)
		return err
	}

	log.Println("Email Sent to address: " + recipient)
	log.Println(result)
	return nil
}
