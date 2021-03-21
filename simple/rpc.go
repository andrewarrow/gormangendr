package simple

import (
	"fmt"
	"net"
	"time"

	"github.com/gocardano/go-cardano-client/cbor"
	"github.com/gocardano/go-cardano-client/multiplex"
)

type DialOption struct {
	Thing string
}
type ClientConn struct {
	conn net.Conn
}

func (cc *ClientConn) Write(b []byte) {
	fmt.Println("write bytes", b)
	n, err := cc.conn.Write(b)
	fmt.Println("n", n, err)
	bytesHeader := make([]byte, 8)
	n, err = cc.conn.Read(bytesHeader)
	fmt.Println("n", n, err)
	fmt.Println("bytesHeader", bytesHeader)
	header, _ := multiplex.ParseHeader(bytesHeader)
	fmt.Printf("%+v\n", header)

	buff := make([]byte, 512)
	n, err = cc.conn.Read(buff)
	fmt.Println("n", n, err, buff)
	// 131 1 1 26 45 150 74 9 121 0 4 108 128 4 0 5 132 0 245 0 3
	psduInput := []byte{}
	psduInput = append(psduInput, bytesHeader...)
	psduInput = append(psduInput, buff[:n]...)
	sdus, _ := multiplex.ParseServiceDataUnits(psduInput)
	arr := sdus[0].DataItems()[0].(*cbor.Array)
	for _, a := range arr.V {
		fmt.Printf("%+v\n", a)
	}
}

func Dial(target string, opts ...DialOption) (*ClientConn, string) {
	d := net.Dialer{Timeout: time.Second * 6}
	conn, err := d.Dial("tcp", target)
	if err != nil {
		fmt.Println("dial error:", err)
		return nil, err.Error()
	}
	cc := ClientConn{conn}
	return &cc, ""
}
