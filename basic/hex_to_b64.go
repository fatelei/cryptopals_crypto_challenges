package basic

import (
	"github.com/fatelei/cryptopals_crypto_challenges/util"
)

func HexToB64(hex string) (string, error) {
	// https://cryptopals.com/sets/1/challenges/1
	data, err := util.DecodeHex(hex)
	if err != nil {
		return "", err
	}
	return util.Encodeb64(data), nil
}
