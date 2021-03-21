package main

import (
	"context"
	"fmt"
	"gormangendr/simple"
	"log"
	"time"
)

func ClientConnect(port string) {
	fmt.Println("0")
	conn, err := simple.Dial(port)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	go HandleClient(conn)
}
func HandleClient(conn *simple.ClientConn) {
	nc := simple.NewNodeClient(conn)
	context.WithTimeout(context.Background(), time.Second*100)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	fmt.Println("2", nc)
	/*
		defer cancel()
			r, err := nc.Handshake(ctx, &network.HandshakeRequest{})
			if err != nil {
				log.Fatalf("could not: %v", err)
			}
			fmt.Println("|", r.Version, r.Block0, r.NodeId, r.Signature, r.Nonce)
			fmt.Println("3")*/
}
