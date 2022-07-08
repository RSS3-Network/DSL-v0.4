package mrc20_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/mrc20"
)

func TestName(t *testing.T) {
	t.Log(mrc20.EventHashTransfer)
	t.Log(mrc20.EventHashApproval)
	t.Log(mrc20.EventHashLogTransfer)
}
