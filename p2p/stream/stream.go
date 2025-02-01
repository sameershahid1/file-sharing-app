package stream

import (
	"file-sharing/p2p/handler"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
)

func NewStream(h *handler.Handler) (*host.Host, string, error) {
	host, err := libp2p.New()
	if err != nil {
		return nil, "", err
	}

	currentMultiAddress := host.Addrs()[0].String() + "/p2p/" + host.ID().String()

	host.SetStreamHandler("/chat", h.HandlerSendMessage)
	host.SetStreamHandler("/file/send", h.HandlerSendingFile)
	host.SetStreamHandler("/file/recieve", h.HandlerRecievingFile)

	return &host, currentMultiAddress, nil
}
