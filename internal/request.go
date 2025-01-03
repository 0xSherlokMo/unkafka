package internal

import (
	"encoding/binary"
	"fmt"
)

type KafkaRequest struct {
	Size              uint32
	RequestAPIKey     int16
	RequestAPIVersion int16
	CorrelationID     int32
	ClientID          string
}

func DecodeRequest(buf []byte) (KafkaRequest, error) {
	if len(buf) < 12 {
		// If the buffer is less than 12 bytes, we can't decode the request (basic non variable size fields)
		return KafkaRequest{}, fmt.Errorf("buffer is too short got %d bytes", len(buf))
	}

	request := KafkaRequest{
		Size: binary.BigEndian.Uint32(buf[:4]),
	}

	if len(buf) < int(request.Size) {
		return KafkaRequest{}, fmt.Errorf("inconsistant size from buffer")
	}
	request.RequestAPIKey = int16(binary.BigEndian.Uint16(buf[4:6]))
	request.RequestAPIVersion = int16(binary.BigEndian.Uint16(buf[6:8]))
	request.CorrelationID = int32(binary.BigEndian.Uint32(buf[8:12]))

	return request, nil
}
