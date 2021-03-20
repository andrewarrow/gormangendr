package main

import (
	"fmt"
	"net"
	"time"

	"github.com/gocardano/go-cardano-client/cbor"
	"github.com/gocardano/go-cardano-client/multiplex"
)

func HandshakeMany() {
	list := LoadRelays()
	for i, item := range list {
		fmt.Println(i, item)
		Handshake(item)
	}
}

const (
	handshakeMessagePropose uint8 = 0
)

func handshakeRequest() []cbor.DataItem {

	arr := cbor.NewArray()
	arr.Add(cbor.NewPositiveInteger8(handshakeMessagePropose))
	versionTable := cbor.NewMap()
	arr.Add(versionTable)

	versionTable.Add(cbor.NewPositiveInteger8(1), cbor.NewPositiveInteger(764824073))
	versionTable.Add(cbor.NewPositiveInteger16(32770), cbor.NewPositiveInteger(764824073))
	//versionTable.Add(cbor.NewPositiveInteger16(32771), cbor.NewPositiveInteger(764824073))

	return []cbor.DataItem{arr}
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

	miniProtocol := multiplex.MiniProtocolIDMuxControl
	sdu := multiplex.NewServiceDataUnit(miniProtocol, multiplex.MessageModeInitiator, handshakeRequest())
	fmt.Println(sdu.Bytes())

	/*
		conn.Write([]byte{48, 221, 75, 88, 0, 0, 0, 17, 130, 0, 162, 1, 26, 45, 150, 74, 9, 25, 128, 2, 26, 45, 150, 74, 9})
		fmt.Println("trying3", host)
		tmp := make([]byte, 256)
		n, err := conn.Read(tmp)
		fmt.Println("trying4", host)
		fmt.Println(n, err, tmp)
		conn.Write([]byte{244, 185, 165, 64, 0, 5, 0, 2, 129, 0})

		tmp = make([]byte, 256)
		n, err = conn.Read(tmp)
		fmt.Println("trying5", host)
		fmt.Println(n, err, tmp)*/
}
