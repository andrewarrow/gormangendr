package main

import (
	"context"
	"fmt"
	"gormangendr/network"
	"log"
	"time"

	"google.golang.org/grpc"
)

func ClientConnect(port string) {
	fmt.Println("0")
	conn, err := grpc.Dial(port, grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	fmt.Println(conn)
	defer conn.Close()
	c := network.NewNodeClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	fmt.Println("2")
	defer cancel()
	r, err := c.Handshake(ctx, &network.HandshakeRequest{})
	if err != nil {
		log.Fatalf("could not: %v", err)
	}
	fmt.Println("|", r.Version, r.Block0, r.NodeId, r.Signature, r.Nonce)
	fmt.Println("3")

}
