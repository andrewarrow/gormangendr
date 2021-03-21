package main

import (
	"fmt"
	"gormangendr/simple"
	"log"
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
	fmt.Println("2", nc)
	r, err := nc.Handshake(simple.HandshakeRequest{})
	if err != "" {
		log.Fatalf("could not: %v", err)
	}
	fmt.Println("|", r.Version, r.Block0, r.NodeId, r.Signature, r.Nonce)
	fmt.Println("3")
}
