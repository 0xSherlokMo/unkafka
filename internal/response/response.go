package response

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Headers struct {
	CorrelatioId uint32
	ErrorCode    int16
}

type Default struct {
	Size    uint32
	Headers Headers
}

func (k *Default) Encode() []byte {
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
	err = binary.Write(&buffer, binary.BigEndian, k.Headers.ErrorCode)
	if err != nil {
		fmt.Println("Error encoding error code: ", err.Error())
		os.Exit(1)
	}

	return buffer.Bytes()
}
