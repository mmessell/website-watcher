package outbound

type WatchConfigRepo interface {
	ListAll() []Person
}

type Person struct {
	Email    string
	Websites []Website
}

type Website struct {
	StartKey string
	EndKey   string
	Url      string
}
