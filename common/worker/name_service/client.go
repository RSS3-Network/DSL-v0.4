package name_service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/avvydomains/golang-client/avvy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell/contract/character"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	lenscontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/spaceid"
	spaceidcontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/spaceid/contracts"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/unstoppable_domains"
	udcontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/unstoppable_domains/contracts"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/httpx"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/unstoppabledomains/resolution-go/v2"
	"github.com/unstoppabledomains/resolution-go/v2/namingservice"
	goens "github.com/wealdtech/go-ens/v3"
)

func ReverseResolveAll(input string, resolveAll bool) model.NameServiceResult {
	result := model.NameServiceResult{}
	splits := strings.Split(input, ".")
	var address string

	switch splits[len(splits)-1] {
	case "eth":
		// error here means the address doesn't have a primary ENS, and can be ignored
		address, _ = ResolveENS(input)
		result.ENS = input
	case "csb":
		address, _ = ResolveCrossbell(input)
		result.Crossbell = input
	case "lens":
		address, _ = ResolveLens(input)
		result.Lens = input
	case "bnb":
		address, _ = ResolveSpaceID(input)
		result.SpaceID = input
	case "crypto", "nft", "blockchain", "bitcoin", "coin", "wallet", "888", "dao", "x", "zil":
		address, _ = ResolveUnstoppableDomains(input)
		result.UnstoppableDomains = input
	case "bit":
		address, _ = ResolveBit(input)
		result.Bit = input
	case "avax":
		address, _ = ResolveAvvy(input)
		result.Avvy = input
	default:
		if ValidateEthereumAddress(input) {
			address = input
		}
	}
	if address != "" {
		result.Address = address
		if resolveAll {
			ResolveAll(&result)
		}
	}

	return result
}

func ValidateEthereumAddress(address string) bool {
	if address == "" {
		return false
	}

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func ResolveAll(result *model.NameServiceResult) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if result.ENS == "" {
			result.ENS, _ = ResolveENS(result.Address)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if result.Crossbell == "" {
			result.Crossbell, _ = ResolveCrossbell(result.Address)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if result.Lens == "" {
			result.Lens, _ = ResolveLens(result.Address)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if result.SpaceID == "" {
			result.SpaceID, _ = ResolveSpaceID(result.Address)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if result.UnstoppableDomains == "" {
			result.UnstoppableDomains, _ = ResolveUnstoppableDomains(result.Address)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if result.Bit == "" {
			result.Bit, _ = ResolveBit(result.Address)
		}
	}()

	wg.Wait()
}

func ResolveCrossbell(input string) (string, error) {
	var result string
	ethereumClient, err := ethclientx.Global(protocol.NetworkCrossbell)
	if err != nil {
		return "", fmt.Errorf("failed to connect to crossbell rpc: %s", err)
	}

	characterContract, err := character.NewCharacter(crossbell.AddressCharacter, ethereumClient)
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
		return "", fmt.Errorf("failed to connect to bsc rpc: %s", err)
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

func ResolveUnstoppableDomains(input string) (string, error) {
	var result string
	if strings.Contains(input, ".") {
		unsBuilder := resolution.NewUnsBuilder()
		ethClient, err := ethclientx.Global(protocol.NetworkEthereum)
		if err != nil {
			return "", fmt.Errorf("failed to connect to ethereum rpc: %s", err)
		}

		polygonClient, err := ethclientx.Global(protocol.NetworkPolygon)
		if err != nil {
			return "", fmt.Errorf("failed to connect to polygon rpc: %s", err)
		}

		unsBuilder.SetContractBackend(ethClient)
		unsBuilder.SetL2ContractBackend(polygonClient)

		unsResolution, err := unsBuilder.Build()
		if err != nil {
			return "", fmt.Errorf("failed to build unsResolution: %s", err)
		}

		znsResolution, err := resolution.NewZnsBuilder().Build()
		if err != nil {
			return "", fmt.Errorf("failed to build znsResolution: %s", err)
		}

		namingServices := map[string]resolution.NamingService{namingservice.UNS: unsResolution, namingservice.ZNS: znsResolution}
		namingServiceName, _ := resolution.DetectNamingService(input)
		if namingServices[namingServiceName] != nil {
			resolvedAddress, err := namingServices[namingServiceName].Addr(input, "ETH")
			if err != nil {
				return "", fmt.Errorf("failed to result %s: %s", namingServiceName, err)
			}
			result = strings.ToLower(resolvedAddress)
		}
	} else {
		// polygon seems to be the most compatible network for unstoppable domains
		ethereumClient, err := ethclientx.Global(protocol.NetworkPolygon)
		if err != nil {
			return "", fmt.Errorf("failed to connect to polygon rpc: %s", err)
		}
		reverseRegistry, err := udcontract.NewUnstoppableDomains(unstoppable_domains.AddressUDPolygon, ethereumClient)
		if err != nil {
			return "", fmt.Errorf("failed to initiate an unstoppable domains reverse registry: %w", err)
		}

		tokenURI, err := reverseRegistry.ReverseOf(&bind.CallOpts{}, common.HexToAddress(input))
		if err != nil {
			return "", fmt.Errorf("not a unstoppable domains user: %w", err)
		}

		UDResult := &unstoppable_domains.UnstoppableDomain{}
		UDEndpoint := "metadata.unstoppabledomains.com"
		request := http.Request{Method: http.MethodGet, URL: &url.URL{Scheme: "https", Host: UDEndpoint, Path: fmt.Sprintf("/metadata/%s", tokenURI)}}

		err = httpx.DoRequest(context.Background(), http.DefaultClient, &request, &UDResult)
		if err != nil {
			return "", fmt.Errorf("failed to request unstoppabledomains metadata: %s", err)
		}

		result = UDResult.Name
	}

	return result, nil
}

func ResolveBit(input string) (string, error) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer func() {
		cancel()
	}()

	bitResult := &model.BitResult{}
	bitEndpoint := "indexer-v1.did.id"
	request := http.Request{Method: http.MethodPost, URL: &url.URL{Scheme: "https", Host: bitEndpoint, Path: "/"}}

	if strings.HasSuffix(input, ".bit") {
		request.Body = io.NopCloser(strings.NewReader(fmt.Sprintf(`{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "das_accountInfo",
			"params": [
				{
					"account": "%s"
				}
			]
		}`, input)))
	} else {
		request.Body = io.NopCloser(strings.NewReader(fmt.Sprintf(`{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "das_reverseRecord",
			"params": [
				{
					"type": "blockchain",
					"key_info": {
						"coin_type": "60",
						"chain_id": "1",
						"key": "%s"
					}
				}
			]
		}`, input)))
	}

	err := httpx.DoRequest(c, http.DefaultClient, &request, &bitResult)

	defer request.Body.Close()

	if err != nil {
		return "", fmt.Errorf("failed to request .bit resolver: %s", err)
	}

	if bitResult.Result.Error != "" {
		return "", fmt.Errorf(".bit resolver returned an error: %s", bitResult.Result.Error)
	}

	if strings.HasSuffix(input, ".bit") {
		return bitResult.Result.Data.AccountInfo.Address, nil
	} else {
		return bitResult.Result.Data.Account, nil
	}
}

func ResolveAvvy(input string) (string, error) {
	var result string

	chainId, _ := strconv.ParseInt(protocol.NetworkToID(protocol.NetworkAvalanche)[2:], 16, 64)

	client := new(avvy.Client)

	networkUrl, err := ethclientx.GlobalUrl(protocol.NetworkAvalanche)
	if err != nil {
		return "", fmt.Errorf("failed to get %s http url: %s", protocol.NetworkAvalanche, err)
	}
	client.Init(networkUrl, int(chainId))

	value, success := client.ResolveStandard(input, client.RECORDS["EVM"])
	if success {
		result = value
	}

	return strings.ToLower(result), nil
}
