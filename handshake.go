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

	queryNode(conn, multiplex.MiniProtocolIDMuxControl, handshakeRequest())

	chainSyncRequest := cbor.NewArray()
	chainSyncRequest.Add(cbor.NewPositiveInteger(0))
	queryNode(conn, multiplex.MiniProtocolIDChainSyncBlocks, []cbor.DataItem{chainSyncRequest})
}

func queryNode(conn net.Conn, miniProtocol multiplex.MiniProtocol, dataItems []cbor.DataItem) {
	sdu := multiplex.NewServiceDataUnit(miniProtocol, multiplex.MessageModeInitiator, dataItems)
	fmt.Println(sdu.Bytes())
	conn.Write(sdu.Bytes())
	tmp := make([]byte, 39)
	conn.Read(tmp)
	fmt.Println("trying", tmp)
}
