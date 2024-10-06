package network


import (
    "context"
    "testing"
    "time"
)

func TestNewProxy(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	p, err := NewProxy(ctx)

	if err != nil {
		t.Fatalf("Expected no error, but got error: %v", err)
	}

	t.Log(p)
}

func TestStartProxy(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	p, err := NewProxy(ctx)

	if err != nil {
		t.Fatalf("Expected no error, but got error: %v", err)
	}

	t.Log(p)

	go p.Start()

	// Perform some basic checks here, such as get localhost
	time.Sleep(7 * time.Second)

	p.Stop()
	cancel()
}

