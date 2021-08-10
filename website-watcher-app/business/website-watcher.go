package business

import (
	"fmt"
)

type WebsiteWatcher struct {
	repo WebsiteRepo
}

func NewWebsiteWatcher(repo WebsiteRepo) WebsiteWatcher {
	return WebsiteWatcher{repo: repo}
}

func (ww WebsiteWatcher) Run() bool {

	for _, website := range ww.repo.GetAllWebsites() {
		state, exists := ww.repo.GetWebsiteState(website)
		ww.repo.PutWebsiteState(website, "website")

		if exists {
			fmt.Println("website " + website.Url + " has been visited: " + state)
			// current state diff from old state
			// diff:
			// update state (Get current state, and put to s3)
			// send emails
			for _, email := range website.Emails {
				fmt.Println("Send email to: " + email)
			}
		} else {
			fmt.Println("website " + website.Url + " hasn't been visited: " + state)
			// update state (Get current state, and put to s3)
		}
	}

	return true
}
