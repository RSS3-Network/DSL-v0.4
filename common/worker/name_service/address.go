package name_service

import (
	"context"
	"regexp"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

func IsValidAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

// CheckContractOnEVM returns a list of EVM networks where the address is not a contract.
func CheckContractOnEVM(address string) []string {
	var wg sync.WaitGroup
	var result []string
	for _, n := range protocol.EVMNetworks {
		wg.Add(1)
		network := n
		go func() {
			defer wg.Done()
			if !IsContract(address, network) {
				result = append(result, network)
			}
		}()
	}
	wg.Wait()

	return result
}

func IsContract(address string, network string) bool {
	ethClient, err := ethclientx.Global(network)
	if err != nil {
		return false
	}

	bytecode, err := ethClient.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return false
	}

	return len(bytecode) > 0
}
