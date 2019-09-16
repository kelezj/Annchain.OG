// Copyright © 2019 Annchain Authors <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package tx_types

import (
	"encoding/json"
	"fmt"
	"github.com/annchain/OG/common"
	"github.com/annchain/OG/common/math"
	"github.com/annchain/OG/types"
	"github.com/annchain/kyber/v3/util/random"
	"math/rand"
	"sort"
	"strings"

	"github.com/annchain/OG/common/hexutil"
)

//go:generate msgp

//msgp:tuple Sequencer
type Sequencer struct {
	// TODO: need more states in sequencer to differentiate multiple chains
	types.TxBase
	Issuer         *common.Address
	BlsJointSig    hexutil.Bytes
	BlsJointPubKey hexutil.Bytes
	StateRoot      common.Hash
	Treasure       *math.BigInt
	Proposing      bool `msg:"-"` // is the sequencer is proposal ,did't commit yet ,use this flag to avoid bls sig verification failed
}

//msgp:tuple SequencerJson
type SequencerJson struct {
	types.TxBaseJson
	Issuer         *common.Address `json:"issuer"`
	BlsJointSig    hexutil.Bytes   `json:"bls_joint_sig"`
	BlsJointPubKey hexutil.Bytes   `json:"bls_joint_pub_key"`
	Proposing      bool            `msg:"-",json:"-"`
}

func (s *Sequencer) ToSmallCaseJson() ([]byte, error) {
	if s == nil {
		return nil, nil
	}
	j := SequencerJson{
		TxBaseJson:     *s.TxBase.ToSmallCase(),
		Issuer:         s.Issuer,
		BlsJointSig:    s.BlsJointSig,
		BlsJointPubKey: s.BlsJointPubKey,
		Proposing:      s.Proposing,
	}

	return json.Marshal(&j)
}

func (s *Sequencer) String() string {
	if s.GetSender() == nil {
		return fmt.Sprintf("%s-[nil]-%d-Seq", s.TxBase.String(), s.AccountNonce)
	} else {
		return fmt.Sprintf("%s-[%.10s]-%d-Seq", s.TxBase.String(), s.Sender().String(), s.AccountNonce)
	}

}

func (s *Sequencer) GetGuarantee() *math.BigInt {
	return s.Treasure
}

//msgp:tuple BlsSigSet
type BlsSigSet struct {
	PublicKey    []byte
	BlsSignature []byte
}

//msgp:tuple Sequencers
type Sequencers []*Sequencer

func SampleSequencer() *Sequencer {
	from := common.HexToAddress("0x33")
	return &Sequencer{
		TxBase: types.TxBase{
			Height:       12,
			ParentsHash:  common.Hashes{common.HexToHash("0xCCDD"), common.HexToHash("0xEEFF")},
			Type:         types.TxBaseTypeSequencer,
			AccountNonce: 234,
		},
		Issuer: &from,
	}
}

func RandomSequencer() *Sequencer {
	id := uint64(rand.Int63n(1000))
	from := common.RandomAddress()
	r := random.New()
	seq := &Sequencer{
		TxBase: types.TxBase{
			Hash:         common.RandomHash(),
			Height:       id,
			ParentsHash:  common.Hashes{common.RandomHash(), common.RandomHash()},
			Type:         types.TxBaseTypeSequencer,
			AccountNonce: uint64(rand.Int63n(50000)),
			Weight:       uint64(rand.Int31n(2000)),
		},
		Issuer: &from,
	}
	seq.BlsJointSig = make([]byte, 64)
	seq.BlsJointPubKey = make([]byte, 128)

	random.Bytes(seq.BlsJointPubKey, r)
	random.Bytes(seq.BlsJointSig, r)
	return seq
}

func (s *Sequencer) SignatureTargets() []byte {
	w := types.NewBinaryWriter()

	sort.Sort(s.ParentsHash)
	for _, parent := range s.Parents() {
		w.Write(parent.Bytes)
	}
	w.Write(s.BlsJointPubKey)
	w.Write(s.AccountNonce)
	if !types.CanRecoverPubFromSig {
		w.Write(s.Issuer.Bytes)
	}
	w.Write(s.Height)
	w.Write(s.Weight)
	w.Write(s.StateRoot.Bytes)

	return w.Bytes()
}

func (s *Sequencer) Sender() common.Address {
	return *s.Issuer
}

func (s *Sequencer) GetSender() *common.Address {
	return s.Issuer
}

func (s *Sequencer) Parents() common.Hashes {
	return s.ParentsHash
}

func (s *Sequencer) Number() uint64 {
	return s.GetHeight()
}

func (s *Sequencer) Compare(tx types.Txi) bool {
	switch tx := tx.(type) {
	case *Sequencer:
		if s.GetTxHash().Cmp(tx.GetTxHash()) == 0 {
			return true
		}
		return false
	default:
		return false
	}
}

func (s *Sequencer) GetBase() *types.TxBase {
	return &s.TxBase
}

func (s *Sequencer) GetHead() *SequencerHeader {
	return NewSequencerHead(s.GetTxHash(), s.Height)
}

func (s *Sequencer) Dump() string {
	var phashes []string
	for _, p := range s.ParentsHash {
		phashes = append(phashes, p.Hex())
	}
	return fmt.Sprintf("pHash:[%s], Issuer : %s , Height: %d, Weight: %d, nonce : %d , blspub: %s, signatute : %s, pubkey:  %s root: %s",
		strings.Join(phashes, " ,"),
		s.Issuer.Hex(),
		s.Height,
		s.Weight,
		s.AccountNonce,
		s.BlsJointPubKey,
		hexutil.Encode(s.PublicKey),
		hexutil.Encode(s.Signature),
		s.StateRoot.Hex(),
	)
}

func (s *Sequencer) RawSequencer() *RawSequencer {
	if s == nil {
		return nil
	}
	return &RawSequencer{
		TxBase:         s.TxBase,
		BlsJointSig:    s.BlsJointSig,
		BlsJointPubKey: s.BlsJointPubKey,
		StateRoot:      s.StateRoot,
	}
}

func (s Sequencers) String() string {
	var strs []string
	for _, v := range s {
		strs = append(strs, v.String())
	}
	return strings.Join(strs, ", ")
}

func (s Sequencers) ToHeaders() SequencerHeaders {
	if len(s) == 0 {
		return nil
	}
	var headers SequencerHeaders
	for _, v := range s {
		head := NewSequencerHead(v.Hash, v.Height)
		headers = append(headers, head)
	}
	return headers
}

func (seqs Sequencers) ToRawSequencers() RawSequencers {
	if len(seqs) == 0 {
		return nil
	}
	var rawSeqs RawSequencers
	for _, v := range seqs {
		rasSeq := v.RawSequencer()
		rawSeqs = append(rawSeqs, rasSeq)
	}
	return rawSeqs
}

func (s *Sequencer) RawTxi() types.RawTxi {
	return s.RawSequencer()
}

func (s *Sequencer) SetSender(addr common.Address) {
	s.Issuer = &addr
}
