package bft

import (
	"encoding/json"
	"fmt"
	"github.com/annchain/OG/common"
	"github.com/annchain/OG/common/crypto"
	"github.com/annchain/OG/common/hexutil"
	"github.com/annchain/OG/types/msg"
	"time"
)

//go:generate msgp

type Signable interface {
	msg.MsgpMember
	SignatureTargets() []byte
}

//msgp:tuple BftMessage
type BftMessage struct {
	Type    BftMessageType
	Payload Signable
}

func (m *BftMessage) String() string {
	return fmt.Sprintf("%s %+v", m.Type.String(), m.Payload)
}


const (
	TimeoutPropose   = time.Duration(8) * time.Second
	TimeoutPreVote   = time.Duration(8) * time.Second
	TimeoutPreCommit = time.Duration(8) * time.Second
	TimeoutDelta     = time.Duration(1) * time.Second
)

type ValueIdMatchType int

const (
	MatchTypeAny ValueIdMatchType = iota
	MatchTypeByValue
	MatchTypeNil
)

type StepType int

const (
	StepTypePropose StepType = iota
	StepTypePreVote
	StepTypePreCommit
)

func (m StepType) String() string {
	switch m {
	case StepTypePropose:
		return "Proposal"
	case StepTypePreVote:
		return "PreVote"
	case StepTypePreCommit:
		return "PreCommit"
	default:
		return "Unknown"
	}
}

func (m *StepType) MarshalJSON() ([]byte, error) {
	s := m.String()
	return json.Marshal(&s)
}

func (m StepType) IsAfter(o StepType) bool {
	return m > o
}

type ChangeStateEvent struct {
	NewStepType StepType
	HeightRound HeightRound
}

type TendermintContext struct {
	HeightRound HeightRound
	StepType    StepType
}

func (t *TendermintContext) Equal(w WaiterContext) bool {
	v, ok := w.(*TendermintContext)
	if !ok {
		return false
	}
	return t.HeightRound == v.HeightRound && t.StepType == v.StepType
}

func (t *TendermintContext) IsAfter(w WaiterContext) bool {
	v, ok := w.(*TendermintContext)
	if !ok {
		return false
	}
	return t.HeightRound.IsAfter(v.HeightRound) || (t.HeightRound == v.HeightRound && t.StepType.IsAfter(v.StepType))
}

type PeerInfo struct {
	Id             int
	PublicKey      crypto.PublicKey `json:"-"`
	Address        common.Address   `json:"address"`
	PublicKeyBytes hexutil.Bytes    `json:"public_key"`
}

// HeightRoundState is the structure for each Height/Round
// Always keep this state that is higher than current in Partner.States map in order not to miss future things
type HeightRoundState struct {
	MessageProposal                       *MessageProposal // the proposal received in this round
	LockedValue                           Proposal
	LockedRound                           int
	ValidValue                            Proposal
	ValidRound                            int
	Decision                              ConsensusDecision   // final decision of mine in this round
	PreVotes                              []*MessagePreVote   // other peers' PreVotes
	PreCommits                            []*MessagePreCommit // other peers' PreCommits
	Sources                               map[uint16]bool     // for line 55, who send future round so that I may advance?
	StepTypeEqualPreVoteTriggered         bool                // for line 34, FIRST time trigger
	StepTypeEqualOrLargerPreVoteTriggered bool                // for line 36, FIRST time trigger
	StepTypeEqualPreCommitTriggered       bool                // for line 47, FIRST time trigger
	Step                                  StepType            // current step in this round
	StartAt                               time.Time
}

func NewHeightRoundState(total int) *HeightRoundState {
	return &HeightRoundState{
		LockedRound: -1,
		ValidRound:  -1,
		PreVotes:    make([]*MessagePreVote, total),
		PreCommits:  make([]*MessagePreCommit, total),
		Sources:     make(map[uint16]bool),
		StartAt:     time.Now(),
	}
}

type HeightRoundStateMap map[HeightRound]*HeightRoundState

func (h *HeightRoundStateMap) MarshalJSON() ([]byte, error) {
	if h == nil {
		return nil, nil
	}
	m := make(map[string]*HeightRoundState, len(*h))
	for k, v := range *h {
		m[k.String()] = v
	}
	return json.Marshal(&m)
}

// BftStatus records all states of BFT
// consider updating resetStatus() if you want to add things here
type BftStatus struct {
	CurrentHR HeightRound
	N         int // total number of participants
	F         int // max number of Byzantines
	Maj23     int
	Peers     []PeerInfo
	States    HeightRoundStateMap // for line 55, round number -> count
}
