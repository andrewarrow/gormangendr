package simple

type NodeClient struct {
}

func NewNodeClient(conn *ClientConn) *NodeClient {
	nc := NodeClient{}
	return &nc
}
