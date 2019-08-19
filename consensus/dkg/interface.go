package dkg

import (
	"github.com/annchain/OG/common/crypto"
	"github.com/annchain/OG/og/message"
	"github.com/annchain/OG/types/p2p_message"
)

type P2PSender interface {
	BroadcastMessage(messageType message.OGMessageType, message p2p_message.Message)
	AnonymousSendMessage(messageType message.OGMessageType, msg p2p_message.Message, anonymousPubKey *crypto.PublicKey)
	SendToPeer(messageType message.OGMessageType, msg p2p_message.Message, peerId string) error
}