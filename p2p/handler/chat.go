package handler

import (
	"bufio"
	"fmt"

	"github.com/libp2p/go-libp2p/core/network"
)

func (h *Handler) HandlerSendMessage(stream network.Stream) {
	defer stream.Close()
	fmt.Println("New stream from: ", stream.Conn().RemotePeer().String())
	reader := bufio.NewReader(stream)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				stream.Close()
				return
			}
			fmt.Println("Stream closed by peer:", err)
			return
		}

		fmt.Println("Message from peer: ", message)
	}
}
