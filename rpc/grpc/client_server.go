package coregrpc

import (
	"context"
	"net"
	"regexp"

	"google.golang.org/grpc"

	"github.com/cometbft/cometbft/rpc/core"
)

var (
	strippedSchemeRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+.-]*://`)
)

// Config is an gRPC server configuration.
//
// Deprecated: A new gRPC API will be introduced after v0.38.
type Config struct {
	MaxOpenConnections int
}

// StartGRPCServer starts a new gRPC BroadcastAPIServer using the given
// net.Listener.
// NOTE: This function blocks - you may want to call it in a go-routine.
//
// Deprecated: A new gRPC API will be introduced after v0.38.
func StartGRPCServer(env *core.Environment, ln net.Listener) error {
	_ = "STUB: not implemented"
	return nil
}

// Send a keepalive ping every 30s of inactivity.

// Close the connection if the ping is not ACKed within 10s.

// Do not require the client to have an active stream to ping.

// Allow client pings as frequent as every 10s. Clients that
// ping faster will be disconnected.

// blocks until one errors or returns nil

// StartGRPCClient dials the gRPC server using protoAddr and returns a new
// BroadcastAPIClient.
//
// Deprecated: A new gRPC API will be introduced after v0.38.
func StartGRPCClient(protoAddr string) BroadcastAPIClient {
	_ = "STUB: not implemented"
	return *new(BroadcastAPIClient)
}

func dialerFunc(_ context.Context, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// StartBlockAPIGRPCClient dials the gRPC server using protoAddr and returns a new
// BlockAPIClient.
func StartBlockAPIGRPCClient(protoAddr string, opts ...grpc.DialOption) (BlockAPIClient, error) {
	_ = "STUB: not implemented"
	return *new(BlockAPIClient), nil
}

// StartBlobstreamAPIGRPCClient dials the gRPC server using protoAddr and returns a new
// BlobstreamAPIClient.
func StartBlobstreamAPIGRPCClient(protoAddr string, opts ...grpc.DialOption) (BlobstreamAPIClient, error) {
	_ = "STUB: not implemented"
	return *new(BlobstreamAPIClient), nil
}

// CanonicalGRPCAddress parses the protoAddr and returns the address, preserving gRPC-supported
// schemes (dns:// and unix://) while stripping other URI schemes.
// Examples:
//   - dns://host:port -> dns://host:port (preserved)
//   - unix:///path -> unix:///path (preserved)
//   - unix:/path -> unix:/path (preserved)
//   - tcp://host:port -> host:port (stripped)
//   - http://host:port -> host:port (stripped)
func CanonicalGRPCAddress(protoAddr string) string {
	_ = "STUB: not implemented"
	// Check if it starts with gRPC-supported schemes - preserve them
	return ""
}

// Strip other URI schemes
