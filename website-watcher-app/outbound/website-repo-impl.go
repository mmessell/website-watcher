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

func (repo WebsiteRepoImpl) GetAllWebsites() []business.Website {
	bytes, _ := repo.bucket.GetObject(repo.configfile)

	var websites []business.Website
	json.Unmarshal(bytes, &websites)

	return websites
}

func (repo WebsiteRepoImpl) GetWebsiteState(website business.Website) (string, bool) {
	bytes, exists := repo.bucket.GetObject(repo.getKeyWithOutputDir(website))
	return string(bytes), exists
}

func (repo WebsiteRepoImpl) PutWebsiteState(website business.Website, state string) {
	repo.bucket.PutObject(repo.getKeyWithOutputDir(website), state)
}

func (repo WebsiteRepoImpl) getKeyWithOutputDir(website business.Website) string {
	return "output/" + website.GetKey() + ".txt"
}
