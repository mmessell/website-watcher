package business

import (
	"errors"
	"log"
	"strings"
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

	return ww.evaluateWebsites(websites)
}

func (ww WebsiteWatcher) evaluateWebsites(websites []Website) error {
	for _, website := range websites {
		oldState, _ := ww.repo.GetWebsiteState(website)

		if oldState != "" {
			log.Print("Website '" + website.Url + "' has been visited.")
			curState, err := ww.updateCurrentState(website)

			if err == nil && curState != oldState {
				log.Print("State change for website '" + website.Url + "'")
				for _, email := range website.Emails {
					log.Print("Send email to: " + email)
				}
			}
		} else {
			log.Print("Website '" + website.Url + "' hasn't been visited.")
			ww.updateCurrentState(website)
		}
	}

	return nil
}

func (ww WebsiteWatcher) updateCurrentState(website Website) (string, error) {
	log.Print("Trying to update state for website '" + website.Url + "'")

	curState, err := ww.getCurrentWebsiteState(website)
	if err != nil {
		log.Fatal("Couldn't get state for website: " + website.Url)
		log.Fatal(err)
		return *new(string), err
	}

	err = ww.repo.PutWebsiteState(website, curState)

	if err != nil {
		log.Fatal("Couldn't update state for website: " + website.Url)
		log.Fatal(err)
		return *new(string), err
	}

	log.Print("State updated for website '" + website.Url + "'")
	return curState, nil
}

func (ww WebsiteWatcher) getCurrentWebsiteState(website Website) (string, error) {
	websiteContent, err := ww.hc.Get()

	if err != nil {
		log.Fatal("Could not get state for website: " + website.Url)
		return *new(string), err
	}

	contentArray := strings.Split(websiteContent, website.StartKey)
	if len(contentArray) > 1 {
		contentArray = strings.Split(contentArray[1], website.EndKey)
	}

	return contentArray[0], nil
}
