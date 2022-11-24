package polygonstaking

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

type Response struct {
	Summary struct {
		Limit     int    `json:"limit"`
		Offset    int    `json:"offset"`
		SortBy    string `json:"sortBy"`
		Direction string `json:"direction"`
		Total     int    `json:"total"`
		Size      int    `json:"size"`
	} `json:"summary"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Result  []struct {
		Id                          int             `json:"id"`
		Name                        string          `json:"name"`
		Description                 *string         `json:"description"`
		LogoUrl                     *string         `json:"logoUrl"`
		Owner                       common.Address  `json:"owner"`
		Signer                      common.Address  `json:"signer"`
		CommissionPercent           decimal.Decimal `json:"commissionPercent"`
		SignerPublicKey             hexutil.Bytes   `json:"signerPublicKey"`
		SelfStake                   decimal.Decimal `json:"selfStake"`
		DelegatedStake              decimal.Decimal `json:"delegatedStake"`
		IsInAuction                 bool            `json:"isInAuction"`
		AuctionAmount               decimal.Decimal `json:"auctionAmount"`
		ClaimedReward               decimal.Decimal `json:"claimedReward"`
		ActivationEpoch             int             `json:"activationEpoch"`
		TotalStaked                 decimal.Decimal `json:"totalStaked"`
		DeactivationEpoch           int             `json:"deactivationEpoch"`
		JailEndEpoch                int             `json:"jailEndEpoch"`
		Status                      string          `json:"status"`
		ContractAddress             common.Address  `json:"contractAddress"`
		UptimePercent               decimal.Decimal `json:"uptimePercent"`
		DelegationEnabled           bool            `json:"delegationEnabled"`
		MissedLatestCheckpointcount int             `json:"missedLatestCheckpointcount"`
	} `json:"result"`
}
