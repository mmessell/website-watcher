package business

type HttpClient interface {
	Get(url string) (string, error)
}
