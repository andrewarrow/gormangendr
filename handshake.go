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
		if i < 4 {
			continue
		}
		Handshake(item)
		break
	}

	for {
		time.Sleep(time.Second * 1)
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
	ClientConnect(host)
}

func queryNode(conn net.Conn, miniProtocol multiplex.MiniProtocol, dataItems []cbor.DataItem) *multiplex.ServiceDataUnit {
	sdu := multiplex.NewServiceDataUnit(miniProtocol, multiplex.MessageModeInitiator, dataItems)
	conn.Write(sdu.Bytes())
	response := []byte{}
	for {
		bytesHeader := make([]byte, 8)
		conn.Read(bytesHeader)
		header, _ := multiplex.ParseHeader(bytesHeader)

		response = append(response, header.Bytes()...)
		totalReadBytes := 0
		for totalReadBytes < header.PayloadLengthAsInt32() {
			buf := make([]byte, header.PayloadLengthAsInt32()-totalReadBytes)
			readBytes, _ := conn.Read(buf)
			response = append(response, buf[:readBytes]...)
			totalReadBytes += readBytes
		}
		if int(header.PayloadLength()) != multiplex.MaxSDUSize {
			break
		}
	}
	sdus, _ := multiplex.ParseServiceDataUnits(response)
	return sdus[0]
}
