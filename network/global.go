package network

import (
	"fmt"
	"time"
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

func Start() {
	for {
		fmt.Println("socket server")
		time.Sleep(time.Second * 1)
	}
}

//Protocol::Grpc
