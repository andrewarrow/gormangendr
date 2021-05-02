package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Address struct {
	Val string
}

func NewAddress() Address {
	a := Address{}
	b := make([]byte, 16)
	rand.Read(b)
	a.Val = fmt.Sprintf("%X", b)
	return a
}

type Config struct {
	Txs          []Tx
	TrustedPeers []string
	BlockHash    string
}

func NewConfig() *Config {
	c := Config{}
	c.TrustedPeers = []string{}
	return &c
}

func (c *Config) GenesisBlockHash() string {
	return "ABC"
}
func (c *Config) AddTrustedPeer(tp string) {
	c.TrustedPeers = append(c.TrustedPeers, tp)
}
func ConfigWithFunds(txs []Tx) *Config {
	c := Config{}
	c.Txs = txs
	return &c
}

type Tx struct {
	Address string
	Val     int64
}

func NewTx(address string, val int64) Tx {
	tx := Tx{}
	tx.Address = address
	tx.Val = val
	return tx
}

/*

   let transaction_message = jcli
       .transaction_builder(block0_hash)
       .build_transaction_from_utxo(
           &utxo,
           *utxo.associated_fund(),
           &sender,
           *utxo.associated_fund(),
           &reciever,
       );
*/

type Gormangendr struct {
	Config *Config
}

func NewGormangendr() *Gormangendr {
	g := Gormangendr{}
	return &g
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

func TwoNodesCommunication() {
	sender := NewAddress()
	receiver := NewAddress()

	fmt.Println(sender.Val)
	fmt.Println(receiver.Val)

	txs := []Tx{}
	txs = append(txs, NewTx("FROG", 100))
	leaderConfig := ConfigWithFunds(txs)
	fmt.Println(leaderConfig)

	leaderGormangendr := NewGormangendr()
	leaderGormangendr.Config = leaderConfig
	go leaderGormangendr.Start()

	trustedConfig := NewConfig()
	trustedConfig.AddTrustedPeer(leaderGormangendr.TrustedPeer())
	trustedConfig.BlockHash = leaderConfig.GenesisBlockHash()
	fmt.Println(trustedConfig)

	for {
		time.Sleep(time.Second)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("testing...")
	TwoNodesCommunication()
}
