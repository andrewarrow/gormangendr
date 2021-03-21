package network

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type GlobalState struct {
	/*block0_hash: HeaderHash,
	  config: Configuration,
	  stats_counter: StatsCounter,
	  topology: P2pTopology,
	  peers: Peers,
	  keypair: NodeKeyPair,
	  span: Span,
	*/
}

func NewGlobalState() *GlobalState {
	gs := GlobalState{}
	return &gs
}

type ConnectionState struct {
}

func NetworkStart(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &NodeService{})
	/*

			struct InboundSubscriptions {
		    pub peer_address: Address,
		    pub block_events: BlockSubscription,
		    pub fragments: FragmentSubscription,
		    pub gossip: GossipSubscription,
		}

				       .block_subscription(comms.subscribe_to_block_announcements()),
				   grpc_client
				       .fragment_subscription(comms.subscribe_to_fragments()),
				   grpc_client
				       .gossip_subscription(comms.subscribe_to_gossip()),

							 InboundSubscriptions
	*/
	//channels
	//pub client_box: MessageBox<ClientMsg>,
	//pub transaction_box: MessageBox<TransactionMsg>,
	//pub block_box: MessageBox<BlockMsg>,

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

//Protocol::Grpc
