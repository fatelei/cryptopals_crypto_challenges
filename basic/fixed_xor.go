package basic

import (
	"fmt"

	"github.com/fatelei/cryptopals_crypto_challenges/util"
)

func FixedXOR(input string, anotherInput string) (string, error) {
	// https://cryptopals.com/sets/1/challenges/2
	a, err := util.DecodeHex(input)
	if err != nil {
		return "", err
	}

	b, err := util.DecodeHex(anotherInput)
	if err != nil {
		return "", err
	}

	rst := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		rst[i] = a[i] ^ b[i]
	}
	return fmt.Sprintf("%x", rst), nil
}
