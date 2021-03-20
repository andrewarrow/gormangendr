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
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "handshake" {
		Handshake(os.Args[2])
	} else if command == "relays" {
		LoadRelays()
	} else if command == "help" {
		PrintHelp()
	}
}
