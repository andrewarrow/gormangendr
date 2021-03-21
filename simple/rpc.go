package simple

import (
	"fmt"
	"net"
	"time"
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
	cc.conn.Read(bytesHeader)
	fmt.Println("bytesHeader", bytesHeader)
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
