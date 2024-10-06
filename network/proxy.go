package network

import (
	"context"
	"fmt"
	"net/http"

	"github.com/elazarl/goproxy"
	// "github.com/elazarl/goproxy/transport"
)

type Proxy struct {
	ctx context.Context
}

func NewProxy(ctx context.Context) (*Proxy, error) {
	result := Proxy{
		ctx: ctx,
	}


	return &result, nil
}

func (p *Proxy) startProxy() {
	proxy := goproxy.NewProxyHttpServer()
	http.ListenAndServe("8080", proxy)
}

func (p* Proxy) Start() {
	for {
		select {
		case <-p.ctx.Done():
			fmt.Println("Background process stopped")
			break
		default:
			fmt.Println("Running background process...")
			// start the proxy
			p.startProxy()
		}
	}
}

func (p* Proxy) Stop() {
}

