package simple

import "context"

type DialOption struct {
	Thing string
}
type ClientConn struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func DialContext(ctx context.Context, target string, opts ...DialOption) (conn *ClientConn, err error) {
	return nil, nil
}

func Dial(target string, opts ...DialOption) (*ClientConn, error) {
	return DialContext(context.Background(), target, opts...)
}
