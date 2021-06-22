package ecrypto_test

import (
	"bytes"
	"testing"

	"github.com/epes/ecrypto"
)

func Test_B64ToBytes(t *testing.T) {
	tests := []struct {
		b64 string
		b   []byte
	}{
		{
			b64: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
			b:   []byte{73, 39, 109, 32, 107, 105, 108, 108, 105, 110, 103, 32, 121, 111, 117, 114, 32, 98, 114, 97, 105, 110, 32, 108, 105, 107, 101, 32, 97, 32, 112, 111, 105, 115, 111, 110, 111, 117, 115, 32, 109, 117, 115, 104, 114, 111, 111, 109},
		},
	}

	for _, tc := range tests {
		got := ecrypto.MustB64ToBytes(tc.b64)

		if !bytes.Equal(tc.b, got) {
			t.Fatalf("expected: %v, got: %v", tc.b, got)
		}
	}
}

func Test_BytesToB64(t *testing.T) {
	tests := []struct {
		b   []byte
		b64 string
	}{
		{
			b:   []byte{73, 39, 109, 32, 107, 105, 108, 108, 105, 110, 103, 32, 121, 111, 117, 114, 32, 98, 114, 97, 105, 110, 32, 108, 105, 107, 101, 32, 97, 32, 112, 111, 105, 115, 111, 110, 111, 117, 115, 32, 109, 117, 115, 104, 114, 111, 111, 109},
			b64: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
	}

	for _, tc := range tests {
		got := ecrypto.BytesToB64(tc.b)

		if tc.b64 != got {
			t.Fatalf("expected: %s, got: %s", tc.b64, got)
		}
	}
}
