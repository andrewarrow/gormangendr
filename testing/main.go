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
	Txs []Tx
}

func NewConfig() *Config {
	cb := Config{}
	return &cb
}

func (cb *Config) GenesisBlockHash() string {
	return "ABC"
}
func ConfigWithFunds(txs []Tx) *Config {
	cb := Config{}
	cb.Txs = txs
	return &cb
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

	for {
		time.Sleep(time.Second)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("testing...")
	TwoNodesCommunication()
}
