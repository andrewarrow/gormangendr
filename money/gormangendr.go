package money

import (
	"fmt"
	"time"
)

type Gormangendr struct {
	Config *Config
}

func NewGormangendr() *Gormangendr {
	g := Gormangendr{}
	return &g
}

func (g *Gormangendr) SendFragment(tx Tx) {
	fmt.Println("SendFragment", tx)
}
func (g *Gormangendr) TrustedPeer() string {
	return "hi"
}
func (g *Gormangendr) Start() {
	for {
		fmt.Println("hi")
		time.Sleep(time.Second)
	}
}
