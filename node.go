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
}

func NewNode() *Node {
	n := Node{}
	return &n
}

func (n *Node) Bootstrap() {
}
func (n *Node) StartServices() {
	for {
		fmt.Println("hi")
		time.Sleep(time.Second * 1)
	}
}
