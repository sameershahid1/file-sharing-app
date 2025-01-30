package handler

import (
	"fmt"

	"github.com/libp2p/go-libp2p/core/network"
)

func (h *Handler) HandlerSendingFile(stream network.Stream) {
	fmt.Println("send file")
}

func (h *Handler) HandlerRecievingFile(stream network.Stream) {
	fmt.Println("recieve file")
}
