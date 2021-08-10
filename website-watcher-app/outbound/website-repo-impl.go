package outbound

import (
	"encoding/json"
	"github.com/mmessell/website-watcher/business"
)

type WebsiteRepoImpl struct {
	bucket     Bucket
	configfile string
}

func NewWebsiteRepoImpl(bucketname string, regionname string, configfile string) WebsiteRepoImpl {
	bucket := NewBucket(bucketname, regionname)
	return WebsiteRepoImpl{bucket: bucket, configfile: configfile}
}

func (repo WebsiteRepoImpl) GetAllWebsites() ([]business.Website, error) {
	bytes, err := repo.bucket.GetObject(repo.configfile)

	if err != nil {
		return nil, err
	}

	var websites []business.Website
	json.Unmarshal(bytes, &websites)

	return websites, nil
}

func (repo WebsiteRepoImpl) GetWebsiteState(website business.Website) (string, error) {
	bytes, err := repo.bucket.GetObject(repo.getKeyWithOutputDir(website))
	return string(bytes), err
}

func (repo WebsiteRepoImpl) PutWebsiteState(website business.Website, state string) error {
	return repo.bucket.PutObject(repo.getKeyWithOutputDir(website), state)
}

func (repo WebsiteRepoImpl) getKeyWithOutputDir(website business.Website) string {
	return "output/" + website.GetKey() + ".txt"
}
