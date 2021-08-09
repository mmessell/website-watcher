package business

import "strings"

type Website struct {
	StartKey string
	EndKey   string
	Url      string
	Emails   []string
}

func (w Website) GetKey() string {
	var key = strings.ReplaceAll(w.Url, ":", "")
	key = strings.ReplaceAll(key, "/", "")
	key = strings.ReplaceAll(key, ".", "")
	key = strings.ReplaceAll(key, "-", "")
	return key
}
