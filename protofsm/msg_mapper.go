package protofsm

import (
	"github.com/lightningnetwork/lnd/fn"
	"github.com/lightningnetwork/lnd/lnwire"
)

// MsgMapper is used to map incoming wire messages into a FSM event. This is
// useful to decouple the translation of an outside or wire message into an
// event type that can be understood by the FSM.
type MsgMapper[Event any] interface {
	// MapMsg maps a wire message into a FSM event. If the message is not
	// mappable, then an None is returned.
	MapMsg(msg lnwire.Message) fn.Option[Event]
}
