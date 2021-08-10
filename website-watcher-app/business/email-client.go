package business

type EmailClient interface {
	Send(recipient string, subject string, message string) error
}
