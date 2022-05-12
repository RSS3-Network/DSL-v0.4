package indexer

import "github.com/naturalselectionlabs/pregod/common/protocol"

type Worker interface {
	Initialize() error
	Handle(message *protocol.Message) error
}
