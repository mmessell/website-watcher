package outbound

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClientImpl struct {
}

func (hc HttpClientImpl) Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return *new(string), err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return *new(string), err
	}

	return string(body), nil
}
