package money

import (
	"fmt"
	"time"
)

type Gormangendr struct {
	Config  *Config
	Passive bool
	MemPool []Tx
}

func NewGormangendr() *Gormangendr {
	g := Gormangendr{}
	g.MemPool = []Tx{}
	return &g
}

func (g *Gormangendr) SendFragment(tx Tx) {
	fmt.Println("SendFragment", tx)
	for _, p := range g.Config.TrustedPeers {
		fmt.Println(p)
	}
}
func (g *Gormangendr) TrustedPeer() string {
	return "hi"
}
func (g *Gormangendr) Start() {
	//start_process
	//jormungandr.wait_for_bootstrap(verification_mode, timeout),
	//self.starter.on_fail,
	//passive or leader:FromGenesis::Hash, leader:FromGenesis::File
	for {
		fmt.Println("hi")
		time.Sleep(time.Second)
	}
}
