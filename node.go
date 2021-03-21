package main

import (
	"fmt"
	"time"
)

type Node struct {
	// pub block0: blockcfg::Block,
	// pub storage: blockchain::Storage,
	// pub rest_context: Option<rest::ContextLock>,
	// pub services: Services,

	//   blockchain: Blockchain,
	//  blockchain_tip: blockchain::Tip,
	//  block0_hash: HeaderHash,
	//  explorer_db: Option<explorer::ExplorerDB>,
	//  rest_context: Option<rest::ContextLock>,
	//  services: Services,

	Blockchain string
	State      string
	Explorer   string
}

func NewNode() *Node {
	n := Node{}
	go n.StartRestService()
	go n.PrepareBlock0()
	return &n
}

func (n *Node) Bootstrap() {
	n.State = "StartingWorkers"
	n.Explorer = "thing"
}
func (n *Node) PrepareBlock0() {
	fmt.Println("prepare")
	time.Sleep(time.Second * 10)
	//  let (blockchain, blockchain_tip) =
	//  start_up::load_blockchain(block0, storage, cache_capacity, settings.rewards_report_all)
	n.Blockchain = "thing"
}
func (n *Node) StartRestService() {
	for {
		fmt.Println("rest")
		time.Sleep(time.Second * 1)
	}
}
func (n *Node) StartServices() {
	for {
		fmt.Println("hi")
		time.Sleep(time.Second * 1)
	}
}
