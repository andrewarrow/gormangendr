package simple

type HandshakeRequest struct {
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
	return res, ""
}

type NodeClient struct {
}

func NewNodeClient(conn *ClientConn) *NodeClient {
	nc := NodeClient{}
	return &nc
}
