package business

type HttpClient interface {
	Get() (string, error)
}
