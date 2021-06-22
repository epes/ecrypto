package ecrypto_test

import (
	"testing"

	"github.com/epes/ecrypto"
)

func Test_HexToB64(t *testing.T) {
	tests := []struct {
		h   string
		b64 string
	}{
		{
			h:   "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			b64: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
	}

	for _, tc := range tests {
		got := ecrypto.BytesToB64(ecrypto.MustHexToBytes(tc.h))

		if tc.b64 != got {
			t.Fatalf("expected: %s, got: %s", tc.b64, got)
		}
	}
}

func Test_ASCIIToB64(t *testing.T) {
	tests := []struct {
		a   string
		b64 string
	}{
		{
			a:   "",
			b64: "",
		},
		{
			a:   "f",
			b64: "Zg==",
		},
		{
			a:   "fo",
			b64: "Zm8=",
		},
		{
			a:   "foo",
			b64: "Zm9v",
		},
		{
			a:   "foob",
			b64: "Zm9vYg==",
		},
		{
			a:   "fooba",
			b64: "Zm9vYmE=",
		},
		{
			a:   "foobar",
			b64: "Zm9vYmFy",
		},
	}

	for _, tc := range tests {
		got := ecrypto.BytesToB64([]byte((tc.a)))

		if tc.b64 != got {
			t.Fatalf("expected: %s, got: %s", tc.b64, got)
		}
	}
}

func Test_B64ToASCII(t *testing.T) {
	tests := []struct {
		a   string
		b64 string
	}{
		{
			a:   "",
			b64: "",
		},
		{
			a:   "f",
			b64: "Zg==",
		},
		{
			a:   "fo",
			b64: "Zm8=",
		},
		{
			a:   "foo",
			b64: "Zm9v",
		},
		{
			a:   "foob",
			b64: "Zm9vYg==",
		},
		{
			a:   "fooba",
			b64: "Zm9vYmE=",
		},
		{
			a:   "foobar",
			b64: "Zm9vYmFy",
		},
	}

	for _, tc := range tests {
		got := string(ecrypto.MustB64ToBytes(tc.b64))

		if tc.a != got {
			t.Fatalf("expected: %s, got: %s", tc.a, got)
		}
	}
}
