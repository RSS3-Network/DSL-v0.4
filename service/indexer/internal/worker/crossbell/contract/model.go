package contract

import (
	"math/big"
)

type ProfileCreated struct {
	Handle    string
	Timestamp *big.Int
}

type SetHandle struct {
	NewHandle string
}

type LinkAddress struct {
	LinkType   [32]byte
	LinklistId *big.Int
}

type LinkAny struct {
	ToURI      string
	LinkType   [32]byte
	LinklistID *big.Int
}

type LinkLinklist struct {
	LinkType [32]byte
}

type LinkProfile struct {
	LinkType   [32]byte
	LinklistId *big.Int
}

type UnlinkProfile struct {
	LinkType [32]byte
}

type SetProfileUri struct {
	NewUri string
}

type PostNote struct {
	LinkItemType [32]byte
	Data         []byte
}
