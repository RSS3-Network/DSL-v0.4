package trigger

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Trigger interface {
	Name() string
	Networks() []string
	Handle(ctx context.Context, message *protocol.Message) error
}
