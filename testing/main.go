package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Address struct {
	Val string
}

func NewAddress() Address {
	a := Address{}
	b := make([]byte, 16)
	rand.Read(b)
	a.Val = fmt.Sprintf("%X", b)
	return a
}

func TwoNodesCommunication() {
	sender := NewAddress()
	receiver := NewAddress()

	fmt.Println(sender.Val)
	fmt.Println(receiver.Val)

}
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("testing...")
	TwoNodesCommunication()
}
