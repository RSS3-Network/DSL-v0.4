package middlewarex

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
)

func GetParamMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		address := c.Param("address")
		if address != "" {
			if address, err := HandleAddress(address); err != nil {
				return err
			} else {
				c.SetParamValues(address)
			}
		}

		return next(c)
	}
}

func HandleAddress(address string) (string, error) {
	result := address

	splits := strings.Split(address, ".")

	if len(splits) < 2 {
		return strings.ToLower(result), nil
	}

	switch splits[len(splits)-1] {
	case "eth":
		nsResult, err := ens.Resolve(config.ConfigHub.RPC.General.Ethereum.HTTP, address)
		if err != nil {
			return "", fmt.Errorf("failed to resolve ens address %s: %s", address, err)
		}

		result = nsResult.Address
	case "csb":
		ethereumClient, err := ethclient.Dial(config.ConfigHub.RPC.General.Crossbell.HTTP)
		if err != nil {
			return "", fmt.Errorf("failed to connect to crossbell rpc: %s", err)
		}

		characterContract, err := crossbell.NewCharacter(crossbell.AddressCharacter, ethereumClient)
		if err != nil {
			return "", fmt.Errorf("failed to connect to crossbell character contract: %s", err)
		}

		character, err := characterContract.GetCharacterByHandle(&bind.CallOpts{}, splits[0])
		if err != nil {
			return "", fmt.Errorf("failed to get crossbell character by handle: %s", err)
		}

		characterOwner, err := characterContract.OwnerOf(&bind.CallOpts{}, character.CharacterId)
		if err != nil {
			return "", fmt.Errorf("failed to get crossbell character owner: %s", err)
		}

		result = characterOwner.String()
	case "lens":
		ethereumClient, err := ethclient.Dial(config.ConfigHub.RPC.General.Polygon.HTTP)
		if err != nil {
			return "", fmt.Errorf("failed to connect to polygon rpc: %s", err)
		}

		lensHubContract, err := lens.NewHub(lens.AddressHub, ethereumClient)
		if err != nil {
			return "", fmt.Errorf("failed to connect to lens hub contract: %s", err)
		}

		profileID, err := lensHubContract.GetProfileIdByHandle(&bind.CallOpts{}, address)
		if err != nil {
			return "", fmt.Errorf("failed to get lens profile id by handle: %s", err)
		}

		profileOwner, err := lensHubContract.OwnerOf(&bind.CallOpts{}, profileID)
		if err != nil {
			return "", fmt.Errorf("failed to get lens profile owner: %s", err)
		}

		result = profileOwner.String()
	}

	return strings.ToLower(result), nil
}
