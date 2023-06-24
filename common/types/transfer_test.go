package types

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/ysmood/got"
)

func TestValidateTypeMap(t *testing.T) {
	g := got.T(t)

	eq := func(tag string, list []string) {
		g.Helper()
		sort.Strings(validTypeMap[tag])
		sort.Strings(list)
		g.Eq(validTypeMap[tag], list)
	}

	eq(filter.TagTransaction, []string{
		filter.TransactionTransfer,
		filter.TransactionMint,
		filter.TransactionBurn,
		filter.TransactionApproval,
		filter.TransactionMultiSig,
		filter.TransactionBridge,
	})

	eq(filter.TagExchange, []string{
		filter.ExchangeWithdraw,
		filter.ExchangeDeposit,
		filter.ExchangeSwap,
		filter.ExchangeLiquidity,
		filter.ExchangeStaking,
	})

	eq(filter.TagCollectible, []string{
		filter.CollectibleTransfer,
		filter.CollectibleAuction,
		filter.CollectibleTrade,
		filter.CollectibleMint,
		filter.CollectibleBurn,
		filter.CollectibleApproval,
		filter.CollectibleEdit,
		filter.CollectibleMusic,
		filter.CollectiblePoap,
	})

	eq(filter.TagSocial, []string{
		filter.SocialPost,
		filter.SocialRevise,
		filter.SocialComment,
		filter.SocialShare,
		filter.SocialProfile,
		filter.SocialFollow,
		filter.SocialUnfollow,
		filter.SocialMint,
		filter.SocialWiki,
		filter.SocialReward,
		filter.SocialProxy,
	})

	eq(filter.TagDonation, []string{
		filter.DonationDonate,
	})

	eq(filter.TagGovernance, []string{
		filter.GovernancePropose,
		filter.GovernanceVote,
	})

	eq(filter.TagMetaverse, []string{
		filter.MetaverseMint,
		filter.MetaverseTrade,
		filter.MetaverseList,
		filter.MetaverseUnlist,
		filter.MetaverseClaim,
	})
}

func TestAllTagTypeCombinations(t *testing.T) {
	g := got.T(t)

	tags := getTags(g)
	types := getTypes(g)

	reg := regexp.MustCompile(`Tag:\s+filter\.(Tag[a-zA-Z]+),\n[\s]+Type:\s+filter\.([a-zA-Z]+),`)

	g.E(filepath.WalkDir(relPath(""), func(path string, d os.DirEntry, err error) error {
		g.E(err)

		if d.IsDir() {
			return nil
		}

		ms := reg.FindAllStringSubmatch(g.Read(path).String(), -1)

		for _, m := range ms {
			tag := tags[m[1]]
			typ := types[m[2]]

			g.Desc("%s %s", m[1], m[2]).True(CheckTypeValid(tag, typ))
		}

		return nil
	}))
}

func getTags(g got.G) map[string]string {
	reg := regexp.MustCompile(`(Tag[a-zA-z]+)\s+string = "(.+?)"`)
	ms := reg.FindAllStringSubmatch(g.Read(relPath("common/protocol/filter/tag.go")).String(), -1)

	tags := map[string]string{}
	for _, m := range ms {
		tags[m[1]] = m[2]
	}

	return tags
}

func getTypes(g got.G) map[string]string {
	reg := regexp.MustCompile(`([a-zA-z]+)\s+string = "(.+?)"`)
	ms := reg.FindAllStringSubmatch(g.Read(relPath("common/protocol/filter/type.go")).String(), -1)

	tags := map[string]string{}
	for _, m := range ms {
		tags[m[1]] = m[2]
	}

	return tags
}

func relPath(path string) string {
	cmd := exec.Command("go", "env", "GOMOD")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return filepath.Join(filepath.Dir(string(out)), filepath.FromSlash(path))
}
