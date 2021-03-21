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

	Blockchain *Blockchain
	State      string
	Explorer   string
}

func NewNode(genesisBlockHash string) *Node {
	n := Node{}
	go n.PrepareBlock0(genesisBlockHash)
	return &n
}

func (n *Node) Bootstrap() {
	n.State = "StartingBoot"
	go n.StartRestService()
	n.Explorer = "thing"
	for {
		time.Sleep(time.Second * 1)
		if n.State == "BootDone" {
			break
		}
	}
}
func (n *Node) PrepareBlock0(genesisBlockHash string) {
	fmt.Println("prepare")
	time.Sleep(time.Second * 10)
	//  let (blockchain, blockchain_tip) =
	//  start_up::load_blockchain(block0, storage, cache_capacity, settings.rewards_report_all)
	n.Blockchain = NewBlockchain(genesisBlockHash)
	fmt.Println("bc ready")
	n.State = "BootDone"
}
func (n *Node) StartRestService() {
	for {
		fmt.Println("rest")
		time.Sleep(time.Second * 1)
	}
}
func (n *Node) StartServices() {
	n.State = "StartingWorkers"
	for {
		fmt.Println("hi")
		time.Sleep(time.Second * 1)
	}
}
