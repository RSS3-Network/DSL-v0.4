package crawler

type Crawler interface {
	Name() string
	Run() error
}
