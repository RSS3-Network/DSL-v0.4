package party

import (
	"fmt"

	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

func BuildPartyURL(partyType string, partyAddress string) string {
	switch partyType {
	case filter.PartyBid:
		return fmt.Sprintf("https://www.partybid.app/party/%s", partyAddress)
	case filter.PartyCollection, filter.PartyBuy:
		return fmt.Sprintf("https://www.partybid.app/buy/%s", partyAddress)
	default:
		return ""
	}
}

func BuildURL(old []string, new ...string) []string {
	if len(new) == 0 {
		return old
	}

	return append(old, new...)
}
