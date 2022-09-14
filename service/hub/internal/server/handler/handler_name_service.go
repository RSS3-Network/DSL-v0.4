package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	lenscontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/spaceid"
	spaceidcontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/spaceid/contracts"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	goens "github.com/wealdtech/go-ens/v3"
	"go.opentelemetry.io/otel"
)

type NameServiceResult struct {
	ENS       string `json:"ens"`
	Crossbell string `json:"crossbell"`
	Lens      string `json:"lens"`
	SpaceID   string `json:"spaceid"`
	Address   string `json:"address"`
}

func (h *Handler) GetNameResolveFunc(c echo.Context) error {
	go h.apiReport(GetNS, c.Get("API-KEY"))
	tracer := otel.Tracer("GetNameResolveFunc")
	_, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	if len(request.Address) == 0 {
		return AddressIsEmpty(c)
	}

	go h.filterReport(GetNS, request)

	splits := strings.Split(request.Address, ".")

	result := NameServiceResult{}
	var address string

	request.Address = strings.ToLower(request.Address)

	switch splits[len(splits)-1] {
	case "eth":
		// error here means the address doesn't have a primary ENS, and can be ignored
		address, _ = ResolveENS(request.Address)
		result.ENS = request.Address
	case "csb":
		address, _ = ResolveCrossbell(request.Address)
		result.Crossbell = request.Address
	case "lens":
		address, _ = ResolveLens(request.Address)
		result.Lens = request.Address
	case "bnb":
		address, _ = ResolveSpaceID(request.Address)
		result.SpaceID = request.Address
	default:
		if ValidateEthereumAddress(request.Address) {
			address = request.Address
		}
	}

	if address != "" {
		result.Address = address
		ResolveAll(&result)
	} else {
		return AddressIsInvalid(c)
	}
	return c.JSON(http.StatusOK, result)
}

// LensResolveFunc temporary function to resolve Lens for Pinata
func (h *Handler) LensResolveFunc(c echo.Context) error {
	result, err := ResolveLens(c.Param("address"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func ResolveCrossbell(input string) (string, error) {
	var result string
	ethereumClient, err := ethclientx.Global(protocol.NetworkCrossbell)
	if err != nil {
		return "", fmt.Errorf("failed to connect to crossbell rpc: %s", err)
	}

	characterContract, err := crossbell.NewCharacter(crossbell.AddressCharacter, ethereumClient)
	if err != nil {
		return "", fmt.Errorf("failed to connect to crossbell character contracts: %s", err)
	}

	if strings.HasSuffix(input, ".csb") {
		character, err := characterContract.GetCharacterByHandle(&bind.CallOpts{}, strings.TrimSuffix(input, ".csb"))
		if err != nil {
			return "", fmt.Errorf("failed to get crossbell character by handle: %s", err)
		}

		characterOwner, err := characterContract.OwnerOf(&bind.CallOpts{}, character.CharacterId)
		if err != nil {
			return "", fmt.Errorf("failed to get crossbell character owner: %s", err)
		}
		result = characterOwner.String()
	} else {
		characterId, err := characterContract.GetPrimaryCharacterId(&bind.CallOpts{}, common.HexToAddress(input))
		if err != nil {
			return "", fmt.Errorf("failed to get crossbell character by address: %s", err)
		}

		result, err = characterContract.GetHandle(&bind.CallOpts{}, characterId)
		if err != nil {
			return "", fmt.Errorf("failed to get crossbell handle by characterId: %s", err)
		}

		if result != "" && !strings.HasSuffix(result, ".csb") {
			result += ".csb"
		}
	}

	return strings.ToLower(result), nil
}

func ResolveENS(address string) (string, error) {
	result, err := ens.Resolve(address)
	if err != nil {
		return "", fmt.Errorf("failed to resolve ENS for: %s", address)
	}

	return strings.ToLower(result), nil
}

func ResolveLens(input string) (string, error) {
	var result string
	ethereumClient, err := ethclientx.Global(protocol.NetworkPolygon)
	if err != nil {
		return "", fmt.Errorf("failed to connect to polygon rpc: %s", err)
	}

	lensHubContract, err := lenscontract.NewHub(lens.HubProxyContractAddress, ethereumClient)
	if err != nil {
		return "", fmt.Errorf("failed to connect to lens hub contracts: %s", err)
	}

	if strings.HasSuffix(input, ".lens") {
		profileID, err := lensHubContract.GetProfileIdByHandle(&bind.CallOpts{}, input)
		if err != nil {
			return "", fmt.Errorf("failed to get lens profile id by handle: %s", err)
		}

		profileOwner, err := lensHubContract.OwnerOf(&bind.CallOpts{}, profileID)
		if err != nil {
			return "", fmt.Errorf("failed to get lens profile owner: %s", err)
		}

		result = profileOwner.String()
	} else {
		profileID, err := lensHubContract.DefaultProfile(&bind.CallOpts{}, common.HexToAddress(input))
		if err != nil {
			return "", fmt.Errorf("failed to get default lens profile id by address: %s", err)
		}

		result, _ = lensHubContract.GetHandle(&bind.CallOpts{}, profileID)
	}

	return strings.ToLower(result), nil
}

func ResolveSpaceID(input string) (string, error) {
	var result string
	ethereumClient, err := ethclientx.Global(protocol.NetworkBinanceSmartChain)
	if err != nil {
		return "", fmt.Errorf("failed to connect to polygon rpc: %s", err)
	}

	spaceidContract, err := spaceidcontract.NewSpaceid(spaceid.AddressSpaceID, ethereumClient)
	if err != nil {
		return "", fmt.Errorf("failed to new a space id contract: %w", err)
	}

	if strings.HasSuffix(input, ".bnb") {
		namehash, _ := goens.NameHash(input)

		resolver, err := spaceidContract.Resolver(&bind.CallOpts{}, namehash)
		if err != nil {
			return "", fmt.Errorf("failed to get space id resolver: %w", err)
		}

		if resolver == ethereum.AddressGenesis {
			return "", fmt.Errorf("the space id does not have a resolver: %s", input)
		}

		spaceidResolver, err := spaceidcontract.NewResolver(resolver, ethereumClient)
		if err != nil {
			return "", fmt.Errorf("failed to new to space id resolver contract: %w", err)
		}

		profileID, err := spaceidResolver.Addr(&bind.CallOpts{}, namehash)
		if err != nil {
			return "", fmt.Errorf("failed to get Space ID by handle: %s", err)
		}

		result = profileID.String()
	} else {
		namehash, _ := goens.NameHash(fmt.Sprintf("%s.addr.reverse", strings.TrimPrefix(input, "0x")))

		resolver, err := spaceidContract.Resolver(&bind.CallOpts{}, namehash)
		if err != nil {
			return "", fmt.Errorf("failed to get space id resolver: %w", err)
		}

		if resolver == ethereum.AddressGenesis {
			return "", fmt.Errorf("the space id does not have a resolver: %s", input)
		}

		spaceidResolver, err := spaceidcontract.NewResolver(resolver, ethereumClient)
		if err != nil {
			return "", fmt.Errorf("failed to new to space id resolver contract: %w", err)
		}

		if result, err = spaceidResolver.Name(&bind.CallOpts{}, namehash); err != nil {
			return "", fmt.Errorf("failed to get space id handle by address: %w", err)
		}
	}

	return strings.ToLower(result), nil
}

func ResolveAll(result *NameServiceResult) {
	if result.ENS == "" {
		result.ENS, _ = ResolveENS(result.Address)
	}

	if result.Crossbell == "" {
		result.Crossbell, _ = ResolveCrossbell(result.Address)
	}

	if result.Lens == "" {
		result.Lens, _ = ResolveLens(result.Address)
	}

	if result.SpaceID == "" {
		result.SpaceID, _ = ResolveSpaceID(result.Address)
	}
}

func ValidateEthereumAddress(address string) bool {
	if address == "" {
		return false
	}

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
