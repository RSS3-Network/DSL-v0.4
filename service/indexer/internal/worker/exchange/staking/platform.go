package staking

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/rss3"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

var UnsupportedPlatform = errors.New("unsupported platform")

var platformMap = map[common.Address]string{
	rss3.AddressStaking: protocol.PlatformRSS3,
}
