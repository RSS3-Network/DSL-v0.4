package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"go.opentelemetry.io/otel"
)

type NameServiceResult struct {
	ENS       string `json:"ens"`
	Crossbell string `json:"crossbell"`
	Lens      string `json:"lens"`
	Address   string `json:"address"`
}

func (h *Handler) GetNameResolve(c echo.Context) error {
	go h.apiReport(GetNS, c.Get("API-KEY"))
	tracer := otel.Tracer("GetENSResolve")
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

func ResolveCrossbell(input string) (string, error) {
	var result string
	ethereumClient, err := ethclient.Dial(config.ConfigHub.RPC.General.Crossbell.HTTP)
	if err != nil {
		return "", fmt.Errorf("failed to connect to crossbell rpc: %s", err)
	}

	characterContract, err := crossbell.NewCharacter(crossbell.AddressCharacter, ethereumClient)
	if err != nil {
		return "", fmt.Errorf("failed to connect to crossbell character contract: %s", err)
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
	result, err := ens.Resolve(config.ConfigHub.RPC.General.Ethereum.HTTP, address)
	if err != nil {
		return "", fmt.Errorf("failed to resolve ENS for: %s", address)
	}

	return strings.ToLower(result), nil
}

func ResolveLens(input string) (string, error) {
	var result string
	ethereumClient, err := ethclient.Dial(config.ConfigHub.RPC.General.Polygon.HTTP)
	if err != nil {
		return "", fmt.Errorf("failed to connect to polygon rpc: %s", err)
	}

	lensHubContract, err := contract.NewHub(lens.ContractAddress, ethereumClient)
	if err != nil {
		return "", fmt.Errorf("failed to connect to lens hub contract: %s", err)
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
}

func ValidateEthereumAddress(address string) bool {
	if address == "" {
		return false
	}

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
