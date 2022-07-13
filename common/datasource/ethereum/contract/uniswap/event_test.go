package uniswap_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
)

func TestName(t *testing.T) {
	t.Log(uniswap.EventHashSwapV2)
	t.Log(uniswap.EventHashSwapV3)
	t.Log(uniswap.EventHashApproval)
}
