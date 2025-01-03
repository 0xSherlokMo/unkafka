package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/kafka-starter-go/internal"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
)

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

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			os.Exit(1)
		}
		request, err := internal.DecodeRequest(buf[:])
		if err != nil {
			fmt.Println("Error decoding request: ", err.Error())
			os.Exit(1)
		}

		response := response.Default{
			Size: 10,
			Headers: response.Headers{
				CorrelatioId: uint32(request.CorrelationID),
			},
		}

		_, err = conn.Write(response.Encode())
		if err != nil {
			fmt.Println("Error writing response: ", err.Error())
			os.Exit(1)
		}
	}
}
