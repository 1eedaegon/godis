package main

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listner, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listner.Close() }() // finally closed
	t.Logf("bound to %q", listner.Addr())
}
