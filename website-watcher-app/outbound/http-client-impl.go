package outbound

type HttpClientImpl struct {
}

func (hc HttpClientImpl) Get() (string, error) {
	return "TEST STATE", nil
}
