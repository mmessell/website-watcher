package business

import (
	"github.com/mmessell/website-watcher/outbound"
)

type MyResponse struct {
	Message string `json:"Answer:"`
}

type WebsiteWatcher struct {
	Wcr outbound.WatchConfigRepo
}

func NewWebsiteWatcher(wcr outbound.WatchConfigRepo) WebsiteWatcher {
	return WebsiteWatcher{Wcr: wcr}
}

func (ww WebsiteWatcher) Run() (MyResponse, error) {
	return MyResponse{Message: ww.Wcr.ListAll()[0].Email}, nil
}
