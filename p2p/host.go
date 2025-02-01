package p2p_host

import (
	"file-sharing/p2p/handler"
	"file-sharing/p2p/stream"
	"fmt"

	"github.com/libp2p/go-libp2p/core/host"
)

type P2pHost struct {
	Handler *handler.Handler
	Host    *host.Host
}

func NewP2pHost() (*P2pHost, string, error) {
	handler := handler.NewHandler()
	host, hostAddress, err := stream.NewStream(handler)
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}

	newP2p2Host := &P2pHost{
		Handler: handler,
		Host:    host,
	}

	return newP2p2Host, hostAddress, nil
}

func (s *P2pHost) Close() error {
	host := *s.Host
	if err := host.Close(); err != nil {
		return err
	}

	return nil
}
