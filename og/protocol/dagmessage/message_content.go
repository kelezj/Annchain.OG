package dagmessage

import (
	"github.com/annchain/OG/common"
	"github.com/annchain/OG/common/math"
)

// all message types related to specific content of DAG are defined here.
//


type MessageContentTx struct {
	Hash         common.Hashes
	ParentsHash  []common.Hashes
	MineNonce    uint64
	AccountNonce uint64
	From         common.Address
	To           common.Address
	Value        *math.BigInt
	TokenId      int32
	PublicKey    []byte
	Data         []byte
	Signature    []byte
}

type MessageContextSequencer struct {
	Hash         common.Hashes
	ParentsHash  []common.Hashes
	MineNonce    uint64
	AccountNonce uint64
	Issuer       common.Address
	PublicKey    []byte
	Signature    []byte
	StateRoot    common.Hash
}
