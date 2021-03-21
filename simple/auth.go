package simple

type ClientAuthRequest struct {
	Signature []byte
	Nonce     []byte
}

func (hr *ClientAuthRequest) Bytes() []byte {
	b := []byte{48, 221, 75, 88, 0, 0, 0, 17, 130, 0, 162, 1, 26, 45, 150, 74, 9, 25, 128, 2, 26, 45, 150, 74, 9}
	return b
}

type ClientAuthResponse struct {
}

func (nc *NodeClient) ClientAuth(req ClientAuthRequest) (ClientAuthResponse, string) {
	res := ClientAuthResponse{}
	nc.conn.Write("ClientAuthResponse", req.Bytes())
	return res, ""
}
