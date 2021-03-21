package simple

type HandshakeRequest struct {
	Nonce []byte
}

func (hr *HandshakeRequest) Bytes() []byte {
	b := []byte{48, 221, 75, 88, 0, 0, 0, 17, 130, 0, 162, 1, 26, 45, 150, 74, 9, 25, 128, 2, 26, 45, 150, 74, 9}
	return b
}

type HandshakeResponse struct {
	Version   int
	Block0    string
	NodeId    string
	Signature []byte
	Nonce     []byte
}

func (nc *NodeClient) Handshake(req HandshakeRequest) (HandshakeResponse, string) {
	res := HandshakeResponse{}
	nc.conn.Write(req.Bytes())
	return res, ""
}

type NodeClient struct {
	conn *ClientConn
}

func NewNodeClient(conn *ClientConn) *NodeClient {
	nc := NodeClient{}
	nc.conn = conn
	return &nc
}
