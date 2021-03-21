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
	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := network.NewNodeClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Handshake(ctx, &network.HandshakeRequest{})
	if err != nil {
		log.Fatalf("could not: %v", err)
	}
	fmt.Println(r)

}
