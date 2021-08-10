package business

import (
	"errors"
	"log"
	"strings"
)

type WebsiteWatcher struct {
	repo WebsiteRepo
	hc   HttpClient
	ec   EmailClient
}

func NewWebsiteWatcher(repo WebsiteRepo, hc HttpClient, ec EmailClient) WebsiteWatcher {
	return WebsiteWatcher{repo: repo, hc: hc, ec: ec}
}

func (ww WebsiteWatcher) Run() error {

	websites, err := ww.repo.GetAllWebsites()

	if err != nil {
		log.Fatal(err)
		return errors.New("Was not able to fetch all websites")
	}

	ww.evaluateWebsites(websites)

	return nil
}

func (ww WebsiteWatcher) evaluateWebsites(websites []Website) {
	for _, website := range websites {
		logMsg := "---------- Evaluating '" + website.Url + "' ----------"
		log.Print(logMsg)
		ww.evaluateWebsite(website)
		log.Print(strings.Repeat("-", len(logMsg)))
	}
}

func (ww WebsiteWatcher) evaluateWebsite(website Website) {
	oldState, _ := ww.repo.GetWebsiteState(website)

	if oldState != "" {
		log.Print("Website '" + website.Url + "' has been visited.")
		curState, err := ww.updateCurrentState(website)

		if err == nil && curState != oldState {
			log.Print("State change for website '" + website.Url + "'")
			err = ww.sendEmails(website)

			if err != nil {
				ww.repo.PutWebsiteState(website, oldState)
			}
		}
	} else {
		log.Print("Website '" + website.Url + "' hasn't been visited.")
		ww.updateCurrentState(website)
	}
}

func (ww WebsiteWatcher) sendEmails(website Website) error {
	for _, email := range website.Emails {
		err := ww.ec.Send(email, "New activity", website.Url)
		if err != nil {
			return err
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
	websiteContent, err := ww.hc.Get(website.Url)

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
