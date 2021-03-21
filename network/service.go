package network

import (
	context "context"
	"log"
)

type NodeService struct {
	UnimplementedGreeterServer
	//channels: Channels,
	//global_state: GlobalStateR,
	//span: Span,
}

func (ns *NodeService) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}
