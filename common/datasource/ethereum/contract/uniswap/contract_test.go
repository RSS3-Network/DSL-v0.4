package uniswap_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
)

func TestName(t *testing.T) {
	t.Log(uniswap.EventHashTransfer)
	t.Log(uniswap.EventHashSwapV2)
	t.Log(uniswap.EventHashSwapV3)
	t.Log(uniswap.EventHashApproval)
	t.Log(uniswap.EventHashMintV2)
	t.Log(uniswap.EventHashBurnV2)
	t.Log(uniswap.EventHashMintV3)
	t.Log(uniswap.EventHashBurnV3)
	t.Log(uniswap.EventHashIncreaseLiquidityV3)
	t.Log(uniswap.EventHashDecreaseLiquidityV3)
	t.Log(uniswap.EventHashCollectV3)
}
