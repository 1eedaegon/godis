// 다이얼은 반드시 타임아웃이 되어야한다.
package dialtimeout

import (
	"net"
	"syscall"
	"testing"
	"time"
)

// If error is temporary, Shoud not disconnect
// ref: https://pkg.go.dev/net#Dialer
func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error {
			return &net.DNSError{
				Err:         "Connection timed out",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true, // unix with windows
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}
	return d.Dial(network, address)
}

func TestDialTimeout(t *testing.T) {
	c, err := DialTimeout("tcp", "10.0.0.1:http", 5*time.Second)
	if err != nil {
		c.Close()
		t.Fatal("Connection didn't timed out")
	}
	nErr, ok := err.(net.Error)
	if !ok {
		t.Fatal(err)
	}
	if !nErr.Timeout() {
		t.Fatal("Error is not a timeout")
	}
}
