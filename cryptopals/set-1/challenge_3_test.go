package set_1

import (
	"testing"

	"github.com/epes/ecrypto/crack"

	"github.com/epes/ecrypto"
)

func Test_SingleByteXORCipher(t *testing.T) {
	tests := []struct {
		h   string
		msg string
	}{
		{
			h:   "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			msg: "Cooking MC's like a pound of bacon",
		},
	}

	for _, tc := range tests {
		msg, _, _ := crack.OneByteXOR(ecrypto.MustHexToBytes(tc.h))

		got := string(msg)

		if tc.msg != got {
			t.Fatalf("expected: '%v', got: '%v'", tc.msg, got)
		}
	}
}
