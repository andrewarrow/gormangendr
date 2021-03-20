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
	response := []byte{}
	for {
		fmt.Println("h1", len(response))
		bytesHeader := make([]byte, multiplex.HeaderSize)
		conn.Read(bytesHeader)
		fmt.Println("header", bytesHeader, len(bytesHeader))
		header, err := multiplex.ParseHeader(bytesHeader)
		fmt.Println("header", header, err)

		response = append(response, header.Bytes()...)
		totalReadBytes := 0
		for totalReadBytes < header.PayloadLengthAsInt32() {
			buf := make([]byte, header.PayloadLengthAsInt32()-totalReadBytes)
			fmt.Println("buf", len(buf), header.PayloadLengthAsInt32())
			readBytes, _ := conn.Read(buf)
			response = append(response, buf[:readBytes]...)
			totalReadBytes += readBytes
		}
		fmt.Println("buf2", header.PayloadLength(), multiplex.MaxSDUSize)
		if int(header.PayloadLength()) != multiplex.MaxSDUSize {
			break
		}
	}
	sdus, _ := multiplex.ParseServiceDataUnits(response)
	fmt.Println(sdus[0].Debug())
	fmt.Println("r", len(sdus), response)
}
