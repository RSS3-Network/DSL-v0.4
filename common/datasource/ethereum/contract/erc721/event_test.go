package erc721_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
)

func TestName(t *testing.T) {
	t.Log(erc721.EventHashTransfer)
	t.Log(erc721.EventHashApproval)
	t.Log(erc721.EventHashApprovalForAll)
}
