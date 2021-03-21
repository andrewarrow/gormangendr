package network

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

//channels
//pub client_box: MessageBox<ClientMsg>,
//pub transaction_box: MessageBox<TransactionMsg>,
//pub block_box: MessageBox<BlockMsg>,

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

func Start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &NodeService{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

//Protocol::Grpc
