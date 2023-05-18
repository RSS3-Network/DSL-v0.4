package name_service

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens/namewrapper"
	registrarv1 "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens/registrar_v1"
	registrarv2 "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens/registrar_v2"
	resolver_v1 "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens/resolver_v1"
	resolver_v2 "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens/resolver_v2"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"

	"go.uber.org/zap"
)

func (i *internal) handleENSNameRenewed(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform string) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	ensContract, err := registrarv1.NewRegistrarV1Filterer(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := ensContract.ParseNameRenewed(log)
	if err != nil {
		return nil, err
	}

	nativeToken, err := i.tokenClient.Native(ctx, message.Network)
	if err != nil {
		return nil, err
	}

	nameServiceMetadata, err := i.buildNameServiceMetadata(ctx, message.Network, fmt.Sprintf("%s.eth", event.Name), filter.CollectibleEditRenew, nativeToken, event.Cost, event.Expires, "", "")
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(transaction.AddressFrom),
		AddressTo:       strings.ToLower(log.Address.String()),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleEdit,
		Metadata:        nameServiceMetadata,
		Network:         transaction.Network,
		Platform:        platform,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleENSNameRegistered(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform string) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	ensContract, err := registrarv1.NewRegistrarV1Filterer(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := ensContract.ParseNameRegistered(log)
	if err != nil {
		return nil, err
	}

	tokenId := new(big.Int).SetBytes(event.Label[:])

	metadataRaw, err := i.buildTokenMetadata(ctx, message.Network, tokenId)
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(transaction.AddressFrom),
		AddressTo:       strings.ToLower(log.Address.String()),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleMint,
		Metadata:        metadataRaw,
		Network:         transaction.Network,
		Platform:        platform,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL(ethereum.BuildTokenURL(message.Network, strings.ToLower(ens.ENSContractAddress.String()), tokenId.String()), ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleENSNameRenewedV2(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform string) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	ensContract, err := registrarv2.NewRegistrarV2Filterer(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := ensContract.ParseNameRenewed(log)
	if err != nil {
		return nil, err
	}

	nativeToken, err := i.tokenClient.Native(ctx, message.Network)
	if err != nil {
		return nil, err
	}

	nameServiceMetadata, err := i.buildNameServiceMetadata(ctx, message.Network, fmt.Sprintf("%s.eth", event.Name), filter.CollectibleEditRenew, nativeToken, event.Cost, event.Expires, "", "")
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(transaction.AddressFrom),
		AddressTo:       strings.ToLower(log.Address.String()),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleEdit,
		Metadata:        nameServiceMetadata,
		Network:         transaction.Network,
		Platform:        platform,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleENSNameRegisteredV2(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform string) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	ensContract, err := registrarv2.NewRegistrarV2Filterer(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := ensContract.ParseNameRegistered(log)
	if err != nil {
		return nil, err
	}

	tokenId := new(big.Int).SetBytes(event.Label[:])

	metadataRaw, err := i.buildTokenMetadata(ctx, message.Network, tokenId)
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(transaction.AddressFrom),
		AddressTo:       strings.ToLower(log.Address.String()),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleMint,
		Metadata:        metadataRaw,
		Network:         transaction.Network,
		Platform:        platform,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL(ethereum.BuildTokenURL(message.Network, strings.ToLower(ens.ENSContractAddress.String()), tokenId.String()), ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleENSNameWrapper(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform string) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	ensContract, err := namewrapper.NewNameWrapperFilterer(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := ensContract.ParseNameWrapped(log)
	if err != nil {
		return nil, err
	}

	var name string
	decodeEnsName, err := decodeName(event.Name)
	if err != nil {
		return nil, err
	}

	if len(decodeEnsName) == 2 {
		name = decodeEnsName[1]
	}

	nameServiceMetadata, err := i.buildNameServiceMetadata(ctx, message.Network, name, filter.CollectibleEditWrap, nil, nil, big.NewInt(0).SetUint64(event.Expiry), "", "")
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(transaction.AddressFrom),
		AddressTo:       strings.ToLower(log.Address.String()),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleEdit,
		Metadata:        nameServiceMetadata,
		Network:         transaction.Network,
		Platform:        platform,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleENSTextChangedV2(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform string) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	ensContract, err := resolver_v2.NewResolverV2Filterer(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := ensContract.ParseTextChanged(log)
	if err != nil {
		return nil, err
	}

	resp, err := i.ensClient.GetEnsName(ctx, common.BytesToHash(event.Node[:]).String())
	if err != nil {
		return nil, err
	}

	if len(resp.Domains) == 0 {
		return nil, fmt.Errorf("not found ens name")
	}

	name := resp.Domains[0].Name

	nameServiceMetadata, err := i.buildNameServiceMetadata(ctx, message.Network, name, filter.CollectibleEditText, nil, nil, nil, event.Key, event.Value)
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(transaction.AddressFrom),
		AddressTo:       strings.ToLower(log.Address.String()),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleEdit,
		Metadata:        nameServiceMetadata,
		Network:         transaction.Network,
		Platform:        platform,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleENSTextChangedV1(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform string) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	ensContract, err := resolver_v1.NewResolverV1Filterer(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := ensContract.ParseTextChanged(log)
	if err != nil {
		return nil, err
	}

	resp, err := i.ensClient.GetEnsName(ctx, common.BytesToHash(event.Node[:]).String())
	if err != nil {
		return nil, err
	}

	if len(resp.Domains) == 0 {
		return nil, fmt.Errorf("not found ens name")
	}

	name := resp.Domains[0].Name

	nameServiceMetadata, err := i.buildNameServiceMetadata(ctx, message.Network, name, filter.CollectibleEditText, nil, nil, nil, event.Key, "")
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(transaction.AddressFrom),
		AddressTo:       strings.ToLower(log.Address.String()),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleEdit,
		Metadata:        nameServiceMetadata,
		Network:         transaction.Network,
		Platform:        platform,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

// decodeName 解码 ENS 名称
func decodeName(buf []byte) ([]string, error) {
	offset := 0
	var list []string
	var firstLabel string

	// 检查第一个标签长度是否为 0
	length := int(buf[offset])
	if length == 0 {
		return []string{firstLabel, "."}, nil
	}

	// 解码每个标签
	for length > 0 {
		// 从 buf 中提取标签
		label := string(buf[offset+1 : offset+1+length])

		// 检查标签是否有效
		if !checkValidLabel(label) {
			return nil, errors.New("invalid label")
		}

		// 如果不是第一个标签，则在标签之间添加点号
		if offset > 1 {
			list = append(list, ".")
		} else {
			firstLabel = label
		}

		// 添加标签到列表中
		list = append(list, label)

		// 更新 offset 和 len
		offset += length + 1
		length = int(buf[offset])
	}

	// 返回第一个标签和完整的域名字符串
	return []string{firstLabel, strings.Join(list, "")}, nil
}

// checkValidLabel 检查标签是否有效
func checkValidLabel(name string) bool {
	for i := 0; i < len(name); i++ {
		c := name[i]
		if c == 0 {
			zap.L().Error("Invalid label contained null byte. Skipping.", zap.String("name", name))
			return false
		} else if c == 46 {
			zap.L().Error("Invalid label contained separator char '.'. Skipping.", zap.String("name", name))
			return false
		}
	}

	return true
}
