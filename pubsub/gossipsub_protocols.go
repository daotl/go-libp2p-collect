package pubsub

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/protocol"
)

// WithGossipSubProtocol changes the default protocol of gossipsub
// and make it like the latest '/meshsub/1.1.0'
func WithGossipSubProtocol(proto protocol.ID) pubsub.Option {
	return pubsub.WithGossipSubProtocols(
		[]protocol.ID{proto},
		func(gsf pubsub.GossipSubFeature, i protocol.ID) bool { return true },
	)
}
