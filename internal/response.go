package internal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type ResponseHeaders struct {
	CorrelatioId uint32
}

type KafkaResponse struct {
	Size    uint32
	Headers ResponseHeaders
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
