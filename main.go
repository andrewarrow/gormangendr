package main

import (
	"context"
	"fmt"
	"gormangendr/money"
	"gormangendr/network"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
)

func TwoNodesCommunication() {
	sender := money.NewWallet()
	receiver := money.NewWallet()

	fmt.Println(sender.Id, sender.Value)
	fmt.Println(receiver.Id, receiver.Value)

	adx := []money.Address{}
	adx = append(adx, money.NewAddressWithValue(sender.Id, 100))
	leaderConfig := money.ConfigWithFunds(adx)
	fmt.Println(leaderConfig)

	leaderGormangendr := money.NewGormangendr()
	leaderGormangendr.Config = leaderConfig
	go leaderGormangendr.Start()

	trustedConfig := money.NewConfig()
	trustedConfig.AddTrustedPeer(leaderGormangendr.TrustedPeer())
	trustedConfig.BlockHash = leaderConfig.GenesisBlockHash()
	fmt.Println(trustedConfig)
	trustedGormangendr := money.NewGormangendr()
	trustedGormangendr.Config = trustedConfig
	trustedGormangendr.Passive = true
	go trustedGormangendr.Start()

	utxo := leaderConfig.Block0UtxoForAddress(&sender)
	block0Hash := trustedConfig.GenesisBlockHash()
	fmt.Println(utxo, block0Hash)

	tx := money.NewTx()
	tx.BuildFromHash(block0Hash)
	tx.SetFromUtxo(&utxo, utxo.AssociatedFund(), &sender, utxo.AssociatedFund(), &receiver)
	// Allow the nodes to exchange gossip info before sending
	// the transaction, or it will have no one to send it to
	time.Sleep(time.Second * 5)

	trustedGormangendr.SendFragment(tx)

	for {
		time.Sleep(time.Second)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial("ip:6000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	nc := network.NewNodeClient(conn)
	fmt.Println("nc", nc)
	in := network.HandshakeRequest{}
	hr, err := nc.Handshake(context.Background(), &in)
	fmt.Println("nc", hr, err)

	//TwoNodesCommunication()
}
