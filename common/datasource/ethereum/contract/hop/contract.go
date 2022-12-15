package hop

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressL2AMMWrapperOptimismUSDT = common.HexToAddress("0x7D269D3E0d61A05a0bA976b7DBF8805bF844AF3F") // https://optimistic.etherscan.io/address/0x7D269D3E0d61A05a0bA976b7DBF8805bF844AF3F

	AddressL2AMMWrapperPolygonETH = common.HexToAddress("0xc315239cFb05F1E130E7E28E603CEa4C014c57f0") // https://polygonscan.com/address/0xc315239cfb05f1e130e7e28e603cea4c014c57f0
	AddressL2Bridge               = common.HexToAddress("0xb98454270065A31D71Bf635F6F7Ee6A518dFb849") // https://polygonscan.com/address/0xb98454270065a31d71bf635f6f7ee6a518dfb849

	EventTransferSent = crypto.Keccak256Hash([]byte("TransferSent(bytes32,uint256,address,uint256,bytes32,uint256,uint256,uint256,uint256)"))
)
