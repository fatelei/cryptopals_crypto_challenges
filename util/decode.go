package util

import (
	"encoding/base64"
	"encoding/hex"
)

func DecodeHex(input string) ([]byte, error) {
	data := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(data, []byte(input))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Encodeb64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}
