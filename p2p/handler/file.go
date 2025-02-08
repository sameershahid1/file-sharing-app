package handler

import (
	"bufio"
	"file-sharing/entity"
	"file-sharing/pkg/util"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
)

func (h *Handler) HandlerSendingFile(st network.Stream) {
	defer st.Close()

	chunk := make([]byte, entity.CHUNK_SIZE)
	reader := bufio.NewReader(st)

	//Reading the metda data
	fileName, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println(err, "failed to read the fileName")
		return
	}

	fileName = strings.TrimSpace(fileName)

	fileSize, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println(err, "failed to read the file size")
		return
	}

	//Creating the file
	fileSize = strings.TrimSpace(fileSize)
	timestamp := time.Now().Format("20060102_150405")
	finalFileName := fmt.Sprintf("%s_%s", timestamp, fileName)

	outFile, err := util.CreateFile(entity.STORAGE_DIRECTORY, fileName)
	if err != nil {
		log.Printf("Error creating file %s: %v", entity.STORAGE_DIRECTORY, err)
		return
	}
	defer outFile.Close()

	log.Printf("Final file name: %s", finalFileName)

	//Saving the chunk in the file, by apending it
	totalBytesWritten := 0

	for {
		n, err := reader.Read(chunk)

		if err != nil {
			if err == io.EOF || n == 0 {
				break
			}

			fmt.Println(err, "failed to read the file chunk")
			return
		}

		if n > 0 {
			written, err := outFile.Write(chunk[:n])
			if err != nil {
				log.Printf("Error writing chunk to file: %v", err)
				return
			}
			totalBytesWritten += written
			fmt.Printf("Sent chunk: %d bytes %d\n", len(chunk), n)
		}
	}
}

func (h *Handler) HandlerRecievingFile(stream network.Stream) {
	fmt.Println("recieve file")
}
