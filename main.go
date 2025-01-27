package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	ctx := context.Background()

	host, err := libp2p.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Your Multiaddress:", host.Addrs()[0].String()+"/p2p/"+host.ID().String())
	host.SetStreamHandler("/chat/1.0.0", handlerStream)

	fmt.Println("Enter the multiaddress of a peer to connect (or press Enter to wait):")
	peerAdr := input()
	if peerAdr != "" {
		connectToPeer(ctx, host, peerAdr)
	}

	select {}
}

func connectToPeer(ctx context.Context, host host.Host, peerAdr string) {
	maddr, err := multiaddr.NewMultiaddr(peerAdr)
	if err != nil {
		fmt.Println("Invalid multiaddress:", err)
		return
	}

	info, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		fmt.Println("Failed to parse peer address: ", err)
		return
	}

	err = host.Connect(ctx, *info)
	if err != nil {
		fmt.Println("Failed to connect to peer: ", err)
		return
	}

	fmt.Println("Connected to peer: ", info.ID)

	stream, err := host.NewStream(ctx, info.ID, "/chat/1.0.0")
	if err != nil {
		fmt.Println("Failed to create stream: ", err)
		return
	}
	defer stream.Close()

	writer := bufio.NewWriter(stream)
	fmt.Println("Type messages to send:")
	for {
		message := input()
		if message == "exit" {
			break
		}
		_, err := writer.WriteString(message + "\n")
		if err != nil {
			fmt.Println("Failed to send message:", err)
			break
		}
		writer.Flush()
	}

}

func handlerStream(stream network.Stream) {
	fmt.Println("New stream from: ", stream.Conn().RemotePeer().String())

	reader := bufio.NewReader(stream)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Stream closed by peer:", err)
			return
		}

		fmt.Println("Message from peer: ", message)
	}
}

// Helper function to read user input
func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text[:len(text)-1]
}
