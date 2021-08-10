package business

type WebsiteRepo interface {
	GetAllWebsites() []Website
	GetWebsiteState(website Website) (string, bool)
	PutWebsiteState(website Website, state string)
}
