package business

import "encoding/json"

type MyResponse struct {
	Message string `json:"Answer:"`
}

type WebsiteWatcher struct {
	wcr WatchConfigRepo
}

func NewWebsiteWatcher(wcr WatchConfigRepo) WebsiteWatcher {
	return WebsiteWatcher{wcr: wcr}
}

func (ww WebsiteWatcher) Run() (MyResponse, error) {
	data, _ := json.Marshal(ww.wcr.ListAll())
	return MyResponse{Message: string(data)}, nil
}
