package net

import (
	"net"
)

// Connect dials the given address and returns a net.Conn. The protoAddr argument should be prefixed with the protocol,
// eg. "tcp://127.0.0.1:8080" or "unix:///tmp/test.sock"
func Connect(protoAddr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// ProtocolAndAddress splits an address into the protocol and address components.
// For instance, "tcp://127.0.0.1:8080" will be split into "tcp" and "127.0.0.1:8080".
// If the address has no protocol prefix, the default is "tcp".
func ProtocolAndAddress(listenAddr string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// GetFreePort gets a free port from the operating system.
// Ripped from https://github.com/phayes/freeport.
// BSD-licensed.
func GetFreePort() (int, error) { _ = "STUB: not implemented"; return 0, nil }
