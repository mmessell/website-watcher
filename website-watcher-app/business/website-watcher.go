package business

import (
	"errors"
	"log"
)

type WebsiteWatcher struct {
	repo WebsiteRepo
	hc   HttpClient
}

func NewWebsiteWatcher(repo WebsiteRepo, hc HttpClient) WebsiteWatcher {
	return WebsiteWatcher{repo: repo, hc: hc}
}

func (ww WebsiteWatcher) Run() error {

	websites, err := ww.repo.GetAllWebsites()
	if err != nil {
		log.Fatal(err)
		return errors.New("Was not able to fetch all websites")
	}

	for _, website := range websites {
		oldState, err := ww.repo.GetWebsiteState(website)

		if err == nil {
			log.Print("Website '" + website.Url + "' has been visited.")
			curState, _ := ww.hc.Get()

			if curState != oldState {
				ww.repo.PutWebsiteState(website, curState)

				for _, email := range website.Emails {
					log.Print("Send email to: " + email)
				}
			}
		} else {
			log.Print("Website '" + website.Url + "' hasn't been visited.")
			curState, _ := ww.hc.Get()
			ww.repo.PutWebsiteState(website, curState)
		}
	}

	return nil
}
