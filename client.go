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
	r, err := nc.Handshake()
	if err != "" {
		log.Fatalf("could not: %v", err)
	}
	fmt.Println("|", r.Version, r.ExtraParams)
	fmt.Println("3")

	//nodeId := ReadPublicKey()
	//nc.ClientAuth(nodeId, sig)
	//&request.node_id, &request.signature
}
