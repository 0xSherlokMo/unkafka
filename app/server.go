package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

type Headers struct {
	CorrelatioId uint32
}

type KafkaResponse struct {
	Size    uint32
	Headers Headers
}

func (k *KafkaResponse) Encode() []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, k.Size)
	if err != nil {
		fmt.Println("Error encoding size: ", err.Error())
		os.Exit(1)
	}
	err = binary.Write(&buffer, binary.BigEndian, k.Headers.CorrelatioId)
	if err != nil {
		fmt.Println("Error encoding correlation id: ", err.Error())
		os.Exit(1)
	}

	return buffer.Bytes()
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	response := KafkaResponse{
		Size: 0,
		Headers: Headers{
			CorrelatioId: 7,
		},
	}

	_, err = conn.Write(response.Encode())
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
		os.Exit(1)
	}

}
