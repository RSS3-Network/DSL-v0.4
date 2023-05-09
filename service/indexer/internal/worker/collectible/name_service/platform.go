package name_service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

var platformMap = map[common.Address]string{
	// ENS Registrar Controller
	ens.EnsRegistrarController:   protocol.PlatformEns,
	ens.EnsRegistrarControllerV2: protocol.PlatformEns,
	ens.EnsNameWrapper:           protocol.PlatformEns,
	ens.EnsPublicResolverV1:      protocol.PlatformEns,
	ens.EnsPublicResolverV2:      protocol.PlatformEns,
}
