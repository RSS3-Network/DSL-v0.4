package command

type Interface interface {
	Initialize() error
	Run() error
}
