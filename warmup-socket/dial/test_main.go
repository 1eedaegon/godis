package dial

import (
	"io"
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

func TestDial(t *testing.T) {
	// binding server port
	listner, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		defer func() { done <- struct{}{} }()
		for {
			// Listening on server
			conn, err := listner.Accept()
			if err != nil {
				t.Fatal(err)
			}
			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()
				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Fatal(err)
						}
						return
					}
					t.Logf("received: %q", buf[n:])
				}
			}(conn)
		}
	}()
	conn, err := net.Dial("tcp", listner.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
	<-done
	listner.Close()
	<-done
}
