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
	return &HandshakeResponse{}, nil
}
func (ns *GendrNodeServer) BlockSubscription(Node_BlockSubscriptionServer) error {
	return nil
}
