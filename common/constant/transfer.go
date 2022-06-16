//go:generate stringer -type=TransferTag -linecomment
package constant

// tag
type TransferTag int

// int -> *string
func (i TransferTag) AsInt() *int {
	str := int(i)
	return &str
}

const (
	TransferTagEth      TransferTag = 1  // eth
	TransferTagErc20    TransferTag = 1  // erc20
	TransferTagErc721   TransferTag = 1  // erc721
	TransferTagNFT      TransferTag = 2  // nft
	TransferTagMirror   TransferTag = 3  // mirror
	TransferTagEns      TransferTag = 3  // ens
	TransferTagGitcoin  TransferTag = 3  // gitcoin
	TransferTagPoap     TransferTag = 3  // poap
	TransferTagSwap     TransferTag = 10 // swap
	TransferTagDonation TransferTag = 10 // donation
	TransferTagDomain   TransferTag = 10 // domain
	TransferTagSnapshot TransferTag = 10 // snapshot
)
