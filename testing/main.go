package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Address struct {
	Id    string
	Value int64
}

func NewAddress() Address {
	a := Address{}
	b := make([]byte, 16)
	rand.Read(b)
	a.Id = fmt.Sprintf("%X", b)
	return a
}
func NewAddressWithValue(id string, value int64) Address {
	a := NewAddress()
	a.Id = id
	a.Value = value
	return a
}

func (a *Address) AssociatedFund() string {
	return "AssociatedFund"
}

type Config struct {
	Addresses    []Address
	TrustedPeers []string
	BlockHash    string
}

func NewConfig() *Config {
	c := Config{}
	c.TrustedPeers = []string{}
	return &c
}

func (c *Config) Block0UtxoForAddress(sender *Address) Address {
	a := NewAddress()
	return a
}

func (c *Config) GenesisBlockHash() string {
	return "ABC"
}
func (c *Config) AddTrustedPeer(tp string) {
	c.TrustedPeers = append(c.TrustedPeers, tp)
}
func ConfigWithFunds(adx []Address) *Config {
	c := Config{}
	c.Addresses = adx
	return &c
}

type Tx struct {
	Address string
	Val     int64
}

func NewTx() Tx {
	tx := Tx{}
	return tx
}

func (tx *Tx) BuildFromHash(hash string) {
}
func (tx *Tx) SetFromUtxo(utxo *Address, af1 string,
	sender *Address, af2 string, receiver *Address) {
}

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

	fmt.Println(sender.Id, sender.Value)
	fmt.Println(receiver.Id, receiver.Value)

	adx := []Address{}
	adx = append(adx, NewAddressWithValue(sender.Id, 100))
	leaderConfig := ConfigWithFunds(adx)
	fmt.Println(leaderConfig)

	leaderGormangendr := NewGormangendr()
	leaderGormangendr.Config = leaderConfig
	go leaderGormangendr.Start()

	trustedConfig := NewConfig()
	trustedConfig.AddTrustedPeer(leaderGormangendr.TrustedPeer())
	trustedConfig.BlockHash = leaderConfig.GenesisBlockHash()
	fmt.Println(trustedConfig)
	trustedGormangendr := NewGormangendr()
	trustedGormangendr.Config = trustedConfig
	go trustedGormangendr.Start()

	utxo := leaderConfig.Block0UtxoForAddress(&sender)
	block0Hash := trustedConfig.GenesisBlockHash()
	fmt.Println(utxo, block0Hash)

	tx := NewTx()
	tx.BuildFromHash(block0Hash)
	tx.SetFromUtxo(&utxo, utxo.AssociatedFund(), &sender, utxo.AssociatedFund(), &receiver)

	for {
		time.Sleep(time.Second)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("testing...")
	TwoNodesCommunication()
}
