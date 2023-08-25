package farcaster

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Cast struct {
	Hash                 common.Hash
	ThreadHash           common.Hash
	ParentHash           common.Hash
	AuthorFid            int64
	AuthorUsername       string
	AuthorDisplayname    string
	AuthorPfpUrl         string
	Text                 string
	Media                []string
	PublishedAt          time.Time
	RepliesCount         int64
	ReactionsCount       int64
	RecastsCount         int64
	WatchesCount         int64
	ParentAuthorFid      int64
	ParentAuthorUsername string
}
