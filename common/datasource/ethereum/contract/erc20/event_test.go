package erc20_test

import (
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"testing"
)

func TestName(t *testing.T) {
	t.Log(erc20.EventHashTransfer)
	t.Log(erc20.EventHashApproval)
}
