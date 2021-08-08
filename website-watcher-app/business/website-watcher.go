package business

type MyResponse struct {
	Message string `json:"Answer:"`
}

type WebsiteWatcher struct {
	Wcr WatchConfigRepo
}

func NewWebsiteWatcher(wcr WatchConfigRepo) WebsiteWatcher {
	return WebsiteWatcher{Wcr: wcr}
}

func (ww WebsiteWatcher) Run() (MyResponse, error) {
	return MyResponse{Message: ww.Wcr.ListAll()[0].Email}, nil
}
