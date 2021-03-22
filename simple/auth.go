package simple

type ClientAuthRequest struct {
	// Node ID of the client, the public key of a key pair.
	NodeId []byte
	// Signature of the server's nonce sent in HandshakeResponse,
	// using the private key in the client's key pair.
	Signature []byte
}

func NewClientAuthRequest() ClientAuthRequest {
	car := ClientAuthRequest{}
	return car
}

func (hr *ClientAuthRequest) Bytes() []byte {
	b := []byte{48, 221, 75, 88, 0, 0, 0, 17, 130, 0, 162, 1, 26, 45, 150, 74, 9, 25, 128, 2, 26, 45, 150, 74, 9}
	return b
}

type ClientAuthResponse struct {
}

func (nc *NodeClient) ClientAuth(nodeId, sig []byte) (ClientAuthResponse, string) {
	res := nc.conn.Write("ClientAuthResponse", []byte{})
	return res.(ClientAuthResponse), ""
}
