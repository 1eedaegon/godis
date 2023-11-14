package dialcancelcontext

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

// Early canceled
func TestDialContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	sync := make(chan struct{})

	go func() {
		defer func() { sync <- struct{}{} }()
		var d net.Dialer
		d.Control = func(_, addr string, _ syscall.RawConn) error {
			time.Sleep(time.Second)
			return nil
		}
		conn, err := d.DialContext(ctx, "tcp", "10.0.0.1:80")
		if err != nil {
			t.Log(err)
			return
		}

		conn.Close()
		t.Error("connection didn't timed out!")
	}()
	cancel() // WithCancel()
	<-sync   // End of Conn

	if ctx.Err() != context.Canceled {
		// %q has double quote
		t.Errorf("expected canceled context %q", ctx.Err())
	}
}
