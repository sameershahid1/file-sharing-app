package app

import (
	//"bufio"
	//"encoding/base64"
	"bufio"
	"file-sharing/entity"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/libp2p/go-libp2p/core/peer"
)

func (a *App) FileShare(targetLink string, filePath string) (string, error) {

	//Connecting with the target host
	targetAddress, err := peer.AddrInfoFromString(targetLink)
	if err != nil {
		return "failed to connect with target peer", err
	}

	host := *a.Host
	if err := host.Connect(a.ctx, *targetAddress); err != nil {
		return "failed to connect with target peer", err
	}

	st, err := host.NewStream(a.ctx, targetAddress.ID, "/file/send")
	if err != nil {
		return "failed to connect with target peer", err
	}

	//Sending the file state on the stream
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "failed to get the file info", err
	}

	fileMetaData := fmt.Sprintf("%s\n%d\n", fileInfo.Name(), fileInfo.Size())
	writer := bufio.NewWriter(st)
	_, err = writer.WriteString(fileMetaData)
	if err != nil {
		return "failed to send file", err
	}
	writer.Flush()

	//Opening the file
	file, err := os.Open(filePath)
	if err != nil {
		return "failed to open the file", err
	}
	defer file.Close()

	//Sending the file in chunks
	buf := make([]byte, entity.CHUNK_SIZE)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "failed to send the file to target host", err
		}

		if n > 0 {
			_, err := writer.Write(buf[:n])
			if err != nil {
				return "failed to send file chunk file", err
			}

			writer.Flush()
			log.Printf("Sent chunk: %d bytes\n", len(buf))
		}
	}

	/*
	   totalBytes := len(fileBytes)
	   offset := 0

	   for offset < totalBytes {
	       end := offset + entity.CHUNK_SIZE
	       if end > totalBytes {
	           end = totalBytes
	       }

	       chunk := fileBytes[offset:end]
	       _, err := writer.Write(chunk)
	       if err != nil {
	           return "failed to send file chunk file", err
	       }

	       writer.Flush()
	       log.Printf("Sent chunk: %d bytes (offset %d to %d) of total %d", len(chunk), offset, end, totalBytes)
	       offset = end
	   }
	*/

	if err := st.CloseWrite(); err != nil {
		log.Printf("Warning: failed to close write side of the stream: %s", err)
	}

	msg := "Successfully sended the file"
	return msg, nil
}
