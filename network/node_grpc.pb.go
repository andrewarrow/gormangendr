// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package network

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	// Initial handshake and authentication of the server node.
	Handshake(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*HandshakeResponse, error)
	// Optional authentication of the client node.
	// Called after Handshake.
	ClientAuth(ctx context.Context, in *ClientAuthRequest, opts ...grpc.CallOption) (*ClientAuthResponse, error)
	Tip(ctx context.Context, in *TipRequest, opts ...grpc.CallOption) (*TipResponse, error)
	// Requests for some peers
	Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersResponse, error)
	GetBlocks(ctx context.Context, in *BlockIds, opts ...grpc.CallOption) (Node_GetBlocksClient, error)
	GetHeaders(ctx context.Context, in *BlockIds, opts ...grpc.CallOption) (Node_GetHeadersClient, error)
	GetFragments(ctx context.Context, in *FragmentIds, opts ...grpc.CallOption) (Node_GetFragmentsClient, error)
	// Requests headers of blocks in the chain in the chronological order,
	// given a selection of possible starting blocks known by the requester,
	// and the identifier of the end block to be included in the returned
	// sequence.
	PullHeaders(ctx context.Context, in *PullHeadersRequest, opts ...grpc.CallOption) (Node_PullHeadersClient, error)
	// Requests blocks in the chain in the chronological order, given a selection
	// of possible starting blocks known by the requester, and the identifier of
	// the end block to be included in the returned sequence.
	PullBlocks(ctx context.Context, in *PullBlocksRequest, opts ...grpc.CallOption) (Node_PullBlocksClient, error)
	PullBlocksToTip(ctx context.Context, in *PullBlocksToTipRequest, opts ...grpc.CallOption) (Node_PullBlocksToTipClient, error)
	// Sends headers of blocks to the service in response to a `missing`
	// item received from the BlockSubscription response stream.
	// The headers are streamed the in chronological order of the chain.
	PushHeaders(ctx context.Context, opts ...grpc.CallOption) (Node_PushHeadersClient, error)
	// Uploads blocks to the service in response to a `solicit` item
	// received from the BlockSubscription response stream.
	UploadBlocks(ctx context.Context, opts ...grpc.CallOption) (Node_UploadBlocksClient, error)
	// Establishes a bidirectional stream to exchange information on new
	// blocks created or accepted by the peers.
	BlockSubscription(ctx context.Context, opts ...grpc.CallOption) (Node_BlockSubscriptionClient, error)
	// Establishes a bidirectional stream to exchange information on new
	// block fragments created or accepted by the peers.
	FragmentSubscription(ctx context.Context, opts ...grpc.CallOption) (Node_FragmentSubscriptionClient, error)
	// Establishes a bidirectional stream to exchange information on new
	// network peers.
	GossipSubscription(ctx context.Context, opts ...grpc.CallOption) (Node_GossipSubscriptionClient, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Handshake(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*HandshakeResponse, error) {
	out := new(HandshakeResponse)
	err := c.cc.Invoke(ctx, "/iohk.chain.node.Node/Handshake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ClientAuth(ctx context.Context, in *ClientAuthRequest, opts ...grpc.CallOption) (*ClientAuthResponse, error) {
	out := new(ClientAuthResponse)
	err := c.cc.Invoke(ctx, "/iohk.chain.node.Node/ClientAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Tip(ctx context.Context, in *TipRequest, opts ...grpc.CallOption) (*TipResponse, error) {
	out := new(TipResponse)
	err := c.cc.Invoke(ctx, "/iohk.chain.node.Node/Tip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersResponse, error) {
	out := new(PeersResponse)
	err := c.cc.Invoke(ctx, "/iohk.chain.node.Node/Peers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetBlocks(ctx context.Context, in *BlockIds, opts ...grpc.CallOption) (Node_GetBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[0], "/iohk.chain.node.Node/GetBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeGetBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_GetBlocksClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type nodeGetBlocksClient struct {
	grpc.ClientStream
}

func (x *nodeGetBlocksClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) GetHeaders(ctx context.Context, in *BlockIds, opts ...grpc.CallOption) (Node_GetHeadersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[1], "/iohk.chain.node.Node/GetHeaders", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeGetHeadersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_GetHeadersClient interface {
	Recv() (*Header, error)
	grpc.ClientStream
}

type nodeGetHeadersClient struct {
	grpc.ClientStream
}

func (x *nodeGetHeadersClient) Recv() (*Header, error) {
	m := new(Header)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) GetFragments(ctx context.Context, in *FragmentIds, opts ...grpc.CallOption) (Node_GetFragmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[2], "/iohk.chain.node.Node/GetFragments", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeGetFragmentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_GetFragmentsClient interface {
	Recv() (*Fragment, error)
	grpc.ClientStream
}

type nodeGetFragmentsClient struct {
	grpc.ClientStream
}

func (x *nodeGetFragmentsClient) Recv() (*Fragment, error) {
	m := new(Fragment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) PullHeaders(ctx context.Context, in *PullHeadersRequest, opts ...grpc.CallOption) (Node_PullHeadersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[3], "/iohk.chain.node.Node/PullHeaders", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodePullHeadersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_PullHeadersClient interface {
	Recv() (*Header, error)
	grpc.ClientStream
}

type nodePullHeadersClient struct {
	grpc.ClientStream
}

func (x *nodePullHeadersClient) Recv() (*Header, error) {
	m := new(Header)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) PullBlocks(ctx context.Context, in *PullBlocksRequest, opts ...grpc.CallOption) (Node_PullBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[4], "/iohk.chain.node.Node/PullBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodePullBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_PullBlocksClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type nodePullBlocksClient struct {
	grpc.ClientStream
}

func (x *nodePullBlocksClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) PullBlocksToTip(ctx context.Context, in *PullBlocksToTipRequest, opts ...grpc.CallOption) (Node_PullBlocksToTipClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[5], "/iohk.chain.node.Node/PullBlocksToTip", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodePullBlocksToTipClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_PullBlocksToTipClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type nodePullBlocksToTipClient struct {
	grpc.ClientStream
}

func (x *nodePullBlocksToTipClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) PushHeaders(ctx context.Context, opts ...grpc.CallOption) (Node_PushHeadersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[6], "/iohk.chain.node.Node/PushHeaders", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodePushHeadersClient{stream}
	return x, nil
}

type Node_PushHeadersClient interface {
	Send(*Header) error
	CloseAndRecv() (*PushHeadersResponse, error)
	grpc.ClientStream
}

type nodePushHeadersClient struct {
	grpc.ClientStream
}

func (x *nodePushHeadersClient) Send(m *Header) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodePushHeadersClient) CloseAndRecv() (*PushHeadersResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PushHeadersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) UploadBlocks(ctx context.Context, opts ...grpc.CallOption) (Node_UploadBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[7], "/iohk.chain.node.Node/UploadBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeUploadBlocksClient{stream}
	return x, nil
}

type Node_UploadBlocksClient interface {
	Send(*Block) error
	CloseAndRecv() (*UploadBlocksResponse, error)
	grpc.ClientStream
}

type nodeUploadBlocksClient struct {
	grpc.ClientStream
}

func (x *nodeUploadBlocksClient) Send(m *Block) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeUploadBlocksClient) CloseAndRecv() (*UploadBlocksResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadBlocksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) BlockSubscription(ctx context.Context, opts ...grpc.CallOption) (Node_BlockSubscriptionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[8], "/iohk.chain.node.Node/BlockSubscription", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeBlockSubscriptionClient{stream}
	return x, nil
}

type Node_BlockSubscriptionClient interface {
	Send(*Header) error
	Recv() (*BlockEvent, error)
	grpc.ClientStream
}

type nodeBlockSubscriptionClient struct {
	grpc.ClientStream
}

func (x *nodeBlockSubscriptionClient) Send(m *Header) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeBlockSubscriptionClient) Recv() (*BlockEvent, error) {
	m := new(BlockEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) FragmentSubscription(ctx context.Context, opts ...grpc.CallOption) (Node_FragmentSubscriptionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[9], "/iohk.chain.node.Node/FragmentSubscription", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeFragmentSubscriptionClient{stream}
	return x, nil
}

type Node_FragmentSubscriptionClient interface {
	Send(*Fragment) error
	Recv() (*Fragment, error)
	grpc.ClientStream
}

type nodeFragmentSubscriptionClient struct {
	grpc.ClientStream
}

func (x *nodeFragmentSubscriptionClient) Send(m *Fragment) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeFragmentSubscriptionClient) Recv() (*Fragment, error) {
	m := new(Fragment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) GossipSubscription(ctx context.Context, opts ...grpc.CallOption) (Node_GossipSubscriptionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[10], "/iohk.chain.node.Node/GossipSubscription", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeGossipSubscriptionClient{stream}
	return x, nil
}

type Node_GossipSubscriptionClient interface {
	Send(*Gossip) error
	Recv() (*Gossip, error)
	grpc.ClientStream
}

type nodeGossipSubscriptionClient struct {
	grpc.ClientStream
}

func (x *nodeGossipSubscriptionClient) Send(m *Gossip) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeGossipSubscriptionClient) Recv() (*Gossip, error) {
	m := new(Gossip)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	// Initial handshake and authentication of the server node.
	Handshake(context.Context, *HandshakeRequest) (*HandshakeResponse, error)
	// Optional authentication of the client node.
	// Called after Handshake.
	ClientAuth(context.Context, *ClientAuthRequest) (*ClientAuthResponse, error)
	Tip(context.Context, *TipRequest) (*TipResponse, error)
	// Requests for some peers
	Peers(context.Context, *PeersRequest) (*PeersResponse, error)
	GetBlocks(*BlockIds, Node_GetBlocksServer) error
	GetHeaders(*BlockIds, Node_GetHeadersServer) error
	GetFragments(*FragmentIds, Node_GetFragmentsServer) error
	// Requests headers of blocks in the chain in the chronological order,
	// given a selection of possible starting blocks known by the requester,
	// and the identifier of the end block to be included in the returned
	// sequence.
	PullHeaders(*PullHeadersRequest, Node_PullHeadersServer) error
	// Requests blocks in the chain in the chronological order, given a selection
	// of possible starting blocks known by the requester, and the identifier of
	// the end block to be included in the returned sequence.
	PullBlocks(*PullBlocksRequest, Node_PullBlocksServer) error
	PullBlocksToTip(*PullBlocksToTipRequest, Node_PullBlocksToTipServer) error
	// Sends headers of blocks to the service in response to a `missing`
	// item received from the BlockSubscription response stream.
	// The headers are streamed the in chronological order of the chain.
	PushHeaders(Node_PushHeadersServer) error
	// Uploads blocks to the service in response to a `solicit` item
	// received from the BlockSubscription response stream.
	UploadBlocks(Node_UploadBlocksServer) error
	// Establishes a bidirectional stream to exchange information on new
	// blocks created or accepted by the peers.
	BlockSubscription(Node_BlockSubscriptionServer) error
	// Establishes a bidirectional stream to exchange information on new
	// block fragments created or accepted by the peers.
	FragmentSubscription(Node_FragmentSubscriptionServer) error
	// Establishes a bidirectional stream to exchange information on new
	// network peers.
	GossipSubscription(Node_GossipSubscriptionServer) error
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) Handshake(context.Context, *HandshakeRequest) (*HandshakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}
func (UnimplementedNodeServer) ClientAuth(context.Context, *ClientAuthRequest) (*ClientAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientAuth not implemented")
}
func (UnimplementedNodeServer) Tip(context.Context, *TipRequest) (*TipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tip not implemented")
}
func (UnimplementedNodeServer) Peers(context.Context, *PeersRequest) (*PeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peers not implemented")
}
func (UnimplementedNodeServer) GetBlocks(*BlockIds, Node_GetBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedNodeServer) GetHeaders(*BlockIds, Node_GetHeadersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetHeaders not implemented")
}
func (UnimplementedNodeServer) GetFragments(*FragmentIds, Node_GetFragmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFragments not implemented")
}
func (UnimplementedNodeServer) PullHeaders(*PullHeadersRequest, Node_PullHeadersServer) error {
	return status.Errorf(codes.Unimplemented, "method PullHeaders not implemented")
}
func (UnimplementedNodeServer) PullBlocks(*PullBlocksRequest, Node_PullBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method PullBlocks not implemented")
}
func (UnimplementedNodeServer) PullBlocksToTip(*PullBlocksToTipRequest, Node_PullBlocksToTipServer) error {
	return status.Errorf(codes.Unimplemented, "method PullBlocksToTip not implemented")
}
func (UnimplementedNodeServer) PushHeaders(Node_PushHeadersServer) error {
	return status.Errorf(codes.Unimplemented, "method PushHeaders not implemented")
}
func (UnimplementedNodeServer) UploadBlocks(Node_UploadBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadBlocks not implemented")
}
func (UnimplementedNodeServer) BlockSubscription(Node_BlockSubscriptionServer) error {
	return status.Errorf(codes.Unimplemented, "method BlockSubscription not implemented")
}
func (UnimplementedNodeServer) FragmentSubscription(Node_FragmentSubscriptionServer) error {
	return status.Errorf(codes.Unimplemented, "method FragmentSubscription not implemented")
}
func (UnimplementedNodeServer) GossipSubscription(Node_GossipSubscriptionServer) error {
	return status.Errorf(codes.Unimplemented, "method GossipSubscription not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iohk.chain.node.Node/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Handshake(ctx, req.(*HandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ClientAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ClientAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iohk.chain.node.Node/ClientAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ClientAuth(ctx, req.(*ClientAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Tip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Tip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iohk.chain.node.Node/Tip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Tip(ctx, req.(*TipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Peers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Peers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iohk.chain.node.Node/Peers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Peers(ctx, req.(*PeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockIds)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).GetBlocks(m, &nodeGetBlocksServer{stream})
}

type Node_GetBlocksServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type nodeGetBlocksServer struct {
	grpc.ServerStream
}

func (x *nodeGetBlocksServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_GetHeaders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockIds)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).GetHeaders(m, &nodeGetHeadersServer{stream})
}

type Node_GetHeadersServer interface {
	Send(*Header) error
	grpc.ServerStream
}

type nodeGetHeadersServer struct {
	grpc.ServerStream
}

func (x *nodeGetHeadersServer) Send(m *Header) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_GetFragments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FragmentIds)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).GetFragments(m, &nodeGetFragmentsServer{stream})
}

type Node_GetFragmentsServer interface {
	Send(*Fragment) error
	grpc.ServerStream
}

type nodeGetFragmentsServer struct {
	grpc.ServerStream
}

func (x *nodeGetFragmentsServer) Send(m *Fragment) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_PullHeaders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullHeadersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).PullHeaders(m, &nodePullHeadersServer{stream})
}

type Node_PullHeadersServer interface {
	Send(*Header) error
	grpc.ServerStream
}

type nodePullHeadersServer struct {
	grpc.ServerStream
}

func (x *nodePullHeadersServer) Send(m *Header) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_PullBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullBlocksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).PullBlocks(m, &nodePullBlocksServer{stream})
}

type Node_PullBlocksServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type nodePullBlocksServer struct {
	grpc.ServerStream
}

func (x *nodePullBlocksServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_PullBlocksToTip_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullBlocksToTipRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).PullBlocksToTip(m, &nodePullBlocksToTipServer{stream})
}

type Node_PullBlocksToTipServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type nodePullBlocksToTipServer struct {
	grpc.ServerStream
}

func (x *nodePullBlocksToTipServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_PushHeaders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).PushHeaders(&nodePushHeadersServer{stream})
}

type Node_PushHeadersServer interface {
	SendAndClose(*PushHeadersResponse) error
	Recv() (*Header, error)
	grpc.ServerStream
}

type nodePushHeadersServer struct {
	grpc.ServerStream
}

func (x *nodePushHeadersServer) SendAndClose(m *PushHeadersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodePushHeadersServer) Recv() (*Header, error) {
	m := new(Header)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_UploadBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).UploadBlocks(&nodeUploadBlocksServer{stream})
}

type Node_UploadBlocksServer interface {
	SendAndClose(*UploadBlocksResponse) error
	Recv() (*Block, error)
	grpc.ServerStream
}

type nodeUploadBlocksServer struct {
	grpc.ServerStream
}

func (x *nodeUploadBlocksServer) SendAndClose(m *UploadBlocksResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeUploadBlocksServer) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_BlockSubscription_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).BlockSubscription(&nodeBlockSubscriptionServer{stream})
}

type Node_BlockSubscriptionServer interface {
	Send(*BlockEvent) error
	Recv() (*Header, error)
	grpc.ServerStream
}

type nodeBlockSubscriptionServer struct {
	grpc.ServerStream
}

func (x *nodeBlockSubscriptionServer) Send(m *BlockEvent) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeBlockSubscriptionServer) Recv() (*Header, error) {
	m := new(Header)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_FragmentSubscription_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).FragmentSubscription(&nodeFragmentSubscriptionServer{stream})
}

type Node_FragmentSubscriptionServer interface {
	Send(*Fragment) error
	Recv() (*Fragment, error)
	grpc.ServerStream
}

type nodeFragmentSubscriptionServer struct {
	grpc.ServerStream
}

func (x *nodeFragmentSubscriptionServer) Send(m *Fragment) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeFragmentSubscriptionServer) Recv() (*Fragment, error) {
	m := new(Fragment)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_GossipSubscription_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).GossipSubscription(&nodeGossipSubscriptionServer{stream})
}

type Node_GossipSubscriptionServer interface {
	Send(*Gossip) error
	Recv() (*Gossip, error)
	grpc.ServerStream
}

type nodeGossipSubscriptionServer struct {
	grpc.ServerStream
}

func (x *nodeGossipSubscriptionServer) Send(m *Gossip) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeGossipSubscriptionServer) Recv() (*Gossip, error) {
	m := new(Gossip)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iohk.chain.node.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _Node_Handshake_Handler,
		},
		{
			MethodName: "ClientAuth",
			Handler:    _Node_ClientAuth_Handler,
		},
		{
			MethodName: "Tip",
			Handler:    _Node_Tip_Handler,
		},
		{
			MethodName: "Peers",
			Handler:    _Node_Peers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlocks",
			Handler:       _Node_GetBlocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetHeaders",
			Handler:       _Node_GetHeaders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetFragments",
			Handler:       _Node_GetFragments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullHeaders",
			Handler:       _Node_PullHeaders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullBlocks",
			Handler:       _Node_PullBlocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullBlocksToTip",
			Handler:       _Node_PullBlocksToTip_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PushHeaders",
			Handler:       _Node_PushHeaders_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadBlocks",
			Handler:       _Node_UploadBlocks_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BlockSubscription",
			Handler:       _Node_BlockSubscription_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "FragmentSubscription",
			Handler:       _Node_FragmentSubscription_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GossipSubscription",
			Handler:       _Node_GossipSubscription_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "node.proto",
}
