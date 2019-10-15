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
package protocol_message

import (
	"fmt"
	"github.com/annchain/OG/common"
	"github.com/annchain/OG/common/crypto"
	"github.com/annchain/OG/consensus/campaign"
	"github.com/annchain/OG/consensus/vrf"

	"strings"

	"github.com/annchain/OG/common/math"
)

//go:generate msgp

// compress data ,for p2p  , small size
//msgp:tuple RawTx
type RawTx struct {
	TxBase
	To      common.Address
	Value   *math.BigInt
	Data    []byte
	TokenId int32
}

//msgp:tuple RawActionTx
type RawActionTx struct {
	TxBase
	Action     uint8
	ActionData ActionData
}

//msgp:tuple RawSequencer
type RawSequencer struct {
	TxBase
	BlsJointSig    []byte
	BlsJointPubKey []byte
	StateRoot      common.Hash
}

//msgp:tuple RawCampaign
type RawCampaign struct {
	TxBase
	DkgPublicKey []byte
	Vrf          vrf.VrfInfo
}

//msgp:tuple RawTermChange
type RawTermChange struct {
	TxBase
	TermId uint64
	PkBls  []byte
	SigSet []*SigSet
}

//msgp:tuple RawSequencers
type RawSequencers []*RawSequencer

//msgp:tuple RawCampaigns
type RawCampaigns []*RawCampaign

//msgp:tuple RawTermChanges
type RawTermChanges []*RawTermChange

//msgp:tuple RawActionTxs
type RawActionTxs []*RawActionTx

//msgp:tuple RawTxs
type RawTxs []*RawTx

func (t *RawTx) Tx() *Tx {
	if t == nil {
		return nil
	}
	tx := &Tx{
		TxBase:  t.TxBase,
		To:      t.To,
		Value:   t.Value,
		Data:    t.Data,
		TokenId: t.TokenId,
	}
	if !CanRecoverPubFromSig {
		tx.SetSender(crypto.Signer.AddressFromPubKeyBytes(tx.PublicKey))
	}
	return tx
}

func (t *RawSequencer) Sequencer() *Sequencer {
	if t == nil {
		return nil
	}
	tx := &Sequencer{
		TxBase:         t.TxBase,
		BlsJointPubKey: t.BlsJointPubKey,
		BlsJointSig:    t.BlsJointSig,
		StateRoot:      t.StateRoot,
	}
	if !CanRecoverPubFromSig {
		addr := crypto.Signer.AddressFromPubKeyBytes(tx.PublicKey)
		tx.Issuer = &addr
	}
	return tx
}

func (t *RawActionTx) ActionTx() *ActionTx {
	if t == nil {
		return nil
	}
	if t == nil {
		return nil
	}
	tx := &ActionTx{
		TxBase:     t.TxBase,
		Action:     t.Action,
		ActionData: t.ActionData,
	}

	if !CanRecoverPubFromSig {
		addr := crypto.Signer.AddressFromPubKeyBytes(tx.PublicKey)
		tx.From = &addr
	}
	return tx
}

func (t *RawActionTx) String() string {
	return fmt.Sprintf("%s-%d-rawATX", t.TxBase.String(), t.AccountNonce)
}

func (t *RawTx) String() string {
	return fmt.Sprintf("%s-%d-RawTx", t.TxBase.String(), t.AccountNonce)
}

func (t *RawSequencer) String() string {
	return fmt.Sprintf("%s-%d_%d-RawSeq", t.TxBase.String(), t.AccountNonce, t.Height)
}

func (t *RawTermChange) String() string {
	return fmt.Sprintf("%s-%d_%d-RawTC", t.TxBase.String(), t.AccountNonce, t.Height)
}

func (t *RawCampaign) String() string {
	return fmt.Sprintf("%s-%d_%d-RawCP", t.TxBase.String(), t.AccountNonce, t.Height)
}

func (r RawTxs) Txs() Txs {
	if len(r) == 0 {
		return nil
	}
	var txs Txs
	for _, v := range r {
		tx := v.Tx()
		txs = append(txs, tx)
	}
	return txs
}

func (r RawTxs) Txis() Txis {
	if len(r) == 0 {
		return nil
	}
	var txis Txis
	for _, v := range r {
		tx := v.Tx()
		txis = append(txis, tx)
	}
	return txis
}

func (r RawSequencers) Sequencers() Sequencers {
	if len(r) == 0 {
		return nil
	}
	var seqs Sequencers
	for _, v := range r {
		seq := v.Sequencer()
		seqs = append(seqs, seq)
	}
	return seqs
}

func (r RawSequencers) Txis() Txis {
	if len(r) == 0 {
		return nil
	}
	var txis Txis
	for _, v := range r {
		seq := v.Sequencer()
		txis = append(txis, seq)
	}
	return txis
}

func (seqs RawSequencers) ToHeaders() SequencerHeaders {
	if len(seqs) == 0 {
		return nil
	}
	var headers SequencerHeaders
	for _, v := range seqs {
		head := NewSequencerHead(v.Hash, v.Height)
		headers = append(headers, head)
	}
	return headers
}

func (r RawTxs) String() string {
	var strs []string
	for _, v := range r {
		strs = append(strs, v.String())
	}
	return strings.Join(strs, ", ")
}

func (r RawSequencers) String() string {
	var strs []string
	for _, v := range r {
		strs = append(strs, v.String())
	}
	return strings.Join(strs, ", ")
}

func (rc *RawCampaign) Campaign() *campaign.Campaign {
	if rc == nil {
		return nil
	}
	cp := &campaign.Campaign{
		TxBase:       rc.TxBase,
		DkgPublicKey: rc.DkgPublicKey,
		Vrf:          rc.Vrf,
	}
	if !CanRecoverPubFromSig {
		addr := crypto.Signer.AddressFromPubKeyBytes(rc.PublicKey)
		cp.Issuer = &addr
	}
	return cp
}

func (r *RawTermChange) TermChange() *TermChange {
	if r == nil {
		return nil
	}
	t := &TermChange{
		TxBase: r.TxBase,
		PkBls:  r.PkBls,
		SigSet: r.SigSet,
		TermID: r.TermId,
	}
	if !CanRecoverPubFromSig {
		addr := crypto.Signer.AddressFromPubKeyBytes(r.PublicKey)
		t.Issuer = &addr
	}
	return t
}

func (r RawCampaigns) Campaigns() campaign.Campaigns {
	if len(r) == 0 {
		return nil
	}
	var cs campaign.Campaigns
	for _, v := range r {
		c := v.Campaign()
		cs = append(cs, c)
	}
	return cs
}

func (r RawTermChanges) TermChanges() TermChanges {
	if len(r) == 0 {
		return nil
	}
	var cs TermChanges
	for _, v := range r {
		c := v.TermChange()
		cs = append(cs, c)
	}
	return cs
}

func (r RawTermChanges) String() string {
	var strs []string
	for _, v := range r {
		strs = append(strs, v.String())
	}
	return strings.Join(strs, ", ")
}

func (r RawCampaigns) String() string {
	var strs []string
	for _, v := range r {
		strs = append(strs, v.String())
	}
	return strings.Join(strs, ", ")
}

func (r RawTermChanges) Txis() Txis {
	if len(r) == 0 {
		return nil
	}
	var cs Txis
	for _, v := range r {
		c := v.TermChange()
		cs = append(cs, c)
	}
	return cs
}

func (r RawCampaigns) Txis() Txis {
	if len(r) == 0 {
		return nil
	}
	var cs Txis
	for _, v := range r {
		c := v.Campaign()
		cs = append(cs, c)
	}
	return cs
}

func (r *RawTxs) Len() int {
	if r == nil {
		return 0
	}
	return len(*r)
}

func (r *RawSequencers) Len() int {
	if r == nil {
		return 0
	}
	return len(*r)
}

func (r *RawCampaigns) Len() int {
	if r == nil {
		return 0
	}
	return len(*r)
}

func (r *RawTermChanges) Len() int {
	if r == nil {
		return 0
	}
	return len(*r)
}

type TxisMarshaler []*RawTxMarshaler

func (t *TxisMarshaler) Append(tx Txi) {
	if tx == nil {
		return
	}
	raw := tx.RawTxi()
	if raw == nil {
		return
	}
	m := RawTxMarshaler{raw}
	if t == nil {
		panic("t is nil ")
	}
	*t = append(*t, &m)
}

func (t TxisMarshaler) Len() int {
	if t == nil {
		return 0
	}
	return len(t)
}

func (t TxisMarshaler) String() string {
	var strs []string
	for _, v := range t {
		strs = append(strs, v.String())
	}
	return strings.Join(strs, ", ")
}

func (t TxisMarshaler) Txis() Txis {
	if t == nil {
		return nil
	}
	var txis Txis
	for _, v := range t {
		if v == nil {
			continue
		}
		txis = append(txis, v.Txi())
	}
	return txis
}

func (t *RawTx) Txi() Txi {
	return t.Tx()
}

func (t *RawSequencer) Txi() Txi {
	return t.Sequencer()
}

func (t *RawTermChange) Txi() Txi {
	return t.TermChange()
}

func (t *RawCampaign) Txi() Txi {
	return t.Campaign()
}

func (a *RawActionTx) Txi() Txi {
	return a.ActionTx()
}

//func (t *RawTx) Dump() string  {
//	return t.Tx().Dump()
//}
//
//func (t *RawSequencer) Dump() string {
//	return t.Sequencer().Dump()
//}
//
//func (t *RawTermChange) Dump() string {
//	return t.TermChange().Dump()
//}
//
//func (t *RawCampaign) Dump() string {
//	return t.Campaign().Dump()
//}
//
//func (t*RawCampaign)GetBase() *TxBase{
//	return t.Campaign().GetBase()
//}
//func (t*RawTermChange)GetBase() *TxBase{
//	return t.TermChange().GetBase()
//
//}
//func (t*RawTx)GetBase() *TxBase{
//	return t.Tx().GetBase()
//
//}
//func (t*RawSequencer)GetBase() *TxBase{
//  return t.Sequencer().GetBase()
//}
//
//func (t*RawCampaign)Sender() common.Address{
//	return t.Campaign().Sender()
//}
//func (t*RawTermChange)Sender() common.Address{
//	return t.TermChange().Sender()
//
//}
//func (t*RawTx)Sender() common.Address{
//	return t.Tx().Sender()
//
//}
//func (t*RawSequencer)Sender() common.Address{
//	return t.Sequencer().Sender()
//}
//
//func (t*RawCampaign)SignatureTargets() []byte{
//	return t.Campaign().SignatureTargets()
//}
//func (t*RawTermChange)SignatureTargets() []byte{
//	return t.TermChange().SignatureTargets()
//
//}
//func (t*RawTx)SignatureTargets() []byte{
//	return t.Tx().SignatureTargets()
//
//}
//func (t*RawSequencer)SignatureTargets() []byte{
//	return t.Sequencer().SignatureTargets()
//}