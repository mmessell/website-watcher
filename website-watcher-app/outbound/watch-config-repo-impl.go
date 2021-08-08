package outbound

type WatchConfigRepoImpl struct{}

func (r WatchConfigRepoImpl) ListAll() []Person {
	return []Person{Person{Email: "mmessell@me.com"}}
}
