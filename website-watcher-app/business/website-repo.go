package business

type WebsiteRepo interface {
	GetAllWebsites() ([]Website, error)
	GetWebsiteState(website Website) (string, error)
	PutWebsiteState(website Website, state string) error
}
