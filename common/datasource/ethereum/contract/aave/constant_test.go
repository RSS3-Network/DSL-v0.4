package aave_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/aave"
)

func TestName(t *testing.T) {
	t.Log(aave.EventHashSupplyV2)
	t.Log(aave.EventHashBorrowV2)
	t.Log(aave.EventHashRepayV2)
	t.Log(aave.EventHashWithdrawV2)

	t.Log(aave.EventHashSupplyV3)
	t.Log(aave.EventHashBorrowV3)
	t.Log(aave.EventHashRepayV3)
	t.Log(aave.EventHashWithdrawV3)
}
