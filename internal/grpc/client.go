package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

// GrpcClient is a client for gRPC
type GrpcClient struct {
	conn *grpc.ClientConn
}

// NewGrpcClient creates a new gRPC client
func NewGrpcClient(ctx context.Context, target string) (*GrpcClient, error) {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(_ context.Context, addr string) (net.Conn, error) {
			// ✅ Use the `ctx` passed into NewGrpcClient
			d := net.Dialer{}
			return d.DialContext(ctx, "tcp", addr)
		}),
	}

	conn, err := grpc.NewClient(target, opts...)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}

	return &GrpcClient{
		conn: conn,
	}, nil

}

// GetConn returns the gRPC client connection
func (c *GrpcClient) GetConn() *grpc.ClientConn {
	return c.conn
}

// Close closes the gRPC client connection
func (c *GrpcClient) Close() {
	c.conn.Close()
}
