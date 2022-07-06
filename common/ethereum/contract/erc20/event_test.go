package erc20_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/erc20"
)

func TestName(t *testing.T) {
	t.Log(erc20.EventHashTransfer)
	t.Log(erc20.EventHashApproval)
}
