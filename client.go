package main

import (
	"fmt"
	"gormangendr/simple"
	"log"
)

func ClientConnect(target string) {
	fmt.Println("0")
	conn, err := simple.Dial(target)
	if err != "" {
		log.Fatalf("did not connect: %v", err)
		return
	}
	go HandleClient(conn)
}
func HandleClient(conn *simple.ClientConn) {
	nc := simple.NewNodeClient(conn)
	fmt.Println("2", nc)
	r, err := nc.Handshake(simple.HandshakeRequest{Nonce: []byte{1, 2, 4}})
	if err != "" {
		log.Fatalf("could not: %v", err)
	}
	fmt.Println("|", r.Version, r.ExtraParam)
	fmt.Println("3")
}
