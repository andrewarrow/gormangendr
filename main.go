package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  gormangendr help         # this menu")
	fmt.Println("  gormangendr handshake    # test network connection")
	fmt.Println("  gormangendr relays       # load mainnet relays")
	fmt.Println("  gormangendr connect      #")
	fmt.Println("  gormangendr node         # --genesis-block-hash")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	argMap = argsToMap()
	if command == "handshake" {
		HandshakeMany()
	} else if command == "connect" {
		Handshake(os.Args[2])
	} else if command == "relays" {
		LoadRelays()
	} else if command == "node" {
		EnsureParamPass("genesis-block-hash")
		node := NewNode(argMap["genesis-block-hash"])
		fmt.Println("New Node...")
		node.Bootstrap()
		fmt.Println("Bootstrapped.")
		go node.StartServices()
		fmt.Println("Services running...")
		for {
			time.Sleep(time.Second * 1)
		}
	} else if command == "help" {
		PrintHelp()
	}
}
