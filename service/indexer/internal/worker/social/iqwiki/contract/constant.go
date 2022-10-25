package iqwiki

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressEveripediaSign = common.HexToAddress("0x191a41c307373211D08613B68df4031977589069")
	AddressEveripedia     = common.HexToAddress("0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D")

	EventPosted       = crypto.Keccak256Hash([]byte("Posted(address,string)"))
	EventOwnerUpdated = crypto.Keccak256Hash([]byte("OwnerUpdated(address,address)"))
)
