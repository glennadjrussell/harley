package network

import (
	"context"
	"fmt"
)

type Proxy struct {
	ctx *context.Context
}

func NewProxy(ctx context.Context) (*Proxy, error) {
	result := Proxy{
		ctx: &ctx,
	}


	return &result, nil
}

func (p* Proxy) Start() {
	for {
		select {
		case <-p.ctx.Done():
			fmt.Println("Background procss stopped")
			break
		default:
			fmt.Println("Running background process...")
			// start the proxy
		}
	}
}

func (p* Proxy) Stop() {
}

