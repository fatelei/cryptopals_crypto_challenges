package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixedXOR(t *testing.T) {
	data, err := FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	assert.Nil(t, err)
	assert.Equal(t, "746865206b696420646f6e277420706c6179", data)
}
