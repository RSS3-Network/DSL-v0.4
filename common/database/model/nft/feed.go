package nft

import (
	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/internal/token"
)

type Feed struct {
	Nft         token.NFT       `json:"nft"`
	Actions     []metadata.Post `json:"actions"`
	RelatedUrls pq.StringArray  `json:"related_urls"`
}
