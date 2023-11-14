package dialcontext

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestDialContext(t *testing.T) {
	dl := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	defer cancel()

	var d net.Dialer
	d.Control = func(_, _ string, _ syscall.RawConn) error {
		time.Sleep(5*time.Second + time.Millisecond)
		return nil
	}

	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err != nil {
		conn.Close()
		t.Fatal("connection did not timeout")
	}
	nErr, ok := err.(net.Error)
	if !ok {
		t.Error(err)
	} else {
		if nErr.Timeout() {
			t.Errorf("error is not a timeout")
		}
	}
	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("actually expected deadline exceeded ")
	}
}
