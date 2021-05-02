package network

import (
	context "context"
)

type GendrNodeServer struct {
	UnimplementedNodeServer
	//channels: Channels,
	//global_state: GlobalStateR,
	//span: Span,
}

func (ns *GendrNodeServer) Handshake(ctx context.Context, in *HandshakeRequest) (*HandshakeResponse, error) {
	hr := &HandshakeResponse{}
	hr.Version = uint32(1)
	hr.Block0 = []byte{1, 2, 4}
	hr.NodeId = []byte{5, 7, 8}
	hr.Signature = []byte{3, 6}
	hr.Nonce = []byte{9}
	return hr, nil
}
func (ns *GendrNodeServer) BlockSubscription(Node_BlockSubscriptionServer) error {
	return nil
}
