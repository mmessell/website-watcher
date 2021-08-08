package outbound

import (
	"github.com/mmessell/website-watcher/business"
)

type WatchConfigRepoImpl struct{}

func (r WatchConfigRepoImpl) ListAll() []business.Person {
	return []business.Person{business.Person{Email: "mmessell@me.com"}}
}
