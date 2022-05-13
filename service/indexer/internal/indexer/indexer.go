package indexer

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Worker interface {
	Handle(ctx context.Context, message *protocol.Message) error
}
