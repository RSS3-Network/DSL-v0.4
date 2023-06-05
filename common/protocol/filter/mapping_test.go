package filter

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"testing"

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

	eq(TagTransaction, []string{
		TransactionTransfer,
		TransactionBridge,
		TransactionMint,
		TransactionBurn,
		TransactionApproval,
		TransactionMultiSig,
	})

	eq(TagExchange, []string{
		ExchangeWithdraw,
		ExchangeDeposit,
		ExchangeSwap,
		ExchangeLiquidity,
	})

	eq(TagCollectible, []string{
		CollectibleTransfer,
		CollectibleAuction,
		CollectibleTrade,
		CollectibleMint,
		CollectibleBurn,
		CollectiblePoap,
		CollectibleApproval,
		CollectibleEdit,
		CollectibleEdit,
		CollectibleMusic,
	})

	eq(TagSocial, []string{
		SocialPost,
		SocialRevise,
		SocialComment,
		SocialShare,
		SocialProfile,
		SocialFollow,
		SocialUnfollow,
		SocialLike,
		SocialMint,
		SocialWiki,
		SocialReward,
		SocialProxy,
	})

	eq(TagDonation, []string{
		DonationLaunch,
		DonationDonate,
	})

	eq(TagGovernance, []string{
		GovernancePropose,
		GovernanceVote,
	})

	eq(TagMetaverse, []string{
		MetaverseMint,
		MetaverseTransfer,
		MetaverseTrade,
		MetaverseGift,
		MetaverseList,
		MetaverseUnlist,
		MetaverseClaim,
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

			if inWhiteList(tag, typ) {
				continue
			}

			g.Desc("%s %s", m[1], m[2]).True(CheckTypeValid(tag, typ))
		}

		return nil
	}))
}

var whiteList = []TagTypeCombo{{TagExchange, ExchangeStaking}}

func inWhiteList(tag, typ string) bool {
	for _, combo := range whiteList {
		if combo.Tag == tag && combo.Type == typ {
			return true
		}
	}
	return false
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
