package indexer

import "github.com/naturalselectionlabs/pregod/common/message"

type Worker interface {
	Initialize() error
	Handle(message *message.Message) error
}
