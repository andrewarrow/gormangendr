package main

import (
	"fmt"
	"net"
	"time"
)

func HandshakeMany() {
	list := LoadRelays()
	for i, item := range list {
		fmt.Println(i, item)
		Handshake(item)
	}
}
func Handshake(host string) {
	fmt.Println("trying", host)
	d := net.Dialer{Timeout: time.Second * 6}
	conn, err := d.Dial("tcp", host)
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	fmt.Println("trying2", host)
	defer conn.Close()
	conn.Write([]byte{222, 54, 237, 104, 0, 0, 0, 25, 130, 0, 163, 1, 26, 45, 150, 74, 9, 25, 128, 2, 26, 45, 150, 74, 9, 25, 128, 3, 26, 45, 150, 74, 9})
	fmt.Println("trying3", host)
	tmp := make([]byte, 256)
	n, err := conn.Read(tmp)
	fmt.Println("trying4", host)
	fmt.Println(n, err, tmp)
	conn.Write([]byte{244, 185, 165, 64, 0, 5, 0, 2, 129, 0})

	tmp = make([]byte, 256)
	n, err = conn.Read(tmp)
	fmt.Println("trying5", host)
	fmt.Println(n, err, tmp)
}
