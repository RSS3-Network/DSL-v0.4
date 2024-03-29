package name_service

import (
	"context"
	"regexp"

	"go.opentelemetry.io/otel"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"
)

func IsEvmValidAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func IsEthereumContract(ctx context.Context, address, network string) (bool, error) {
	tracer := otel.Tracer("IsEthereumContract")
	_, httpSnap := tracer.Start(ctx, "name_service")
	defer httpSnap.End()

	ethClient, err := ethclientx.Global(network)
	if err != nil {
		return false, err
	}

	bytecode, err := ethClient.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		loggerx.Global().Error("IsEthereumContract: ethclient CodeAt error: %v, address: %v", zap.Error(err), zap.String("address", address))

		return false, err
	}

	return len(bytecode) > 0, nil
}
