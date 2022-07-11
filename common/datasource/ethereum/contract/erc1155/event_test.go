package erc1155_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc1155"
)

func TestName(t *testing.T) {
	t.Log(erc1155.EventHashTransferSingle)
	t.Log(erc1155.EventHashTransferBatch)
	t.Log(erc1155.EventHashApprovalForAll)
	t.Log(erc1155.EventHashURI)
}
