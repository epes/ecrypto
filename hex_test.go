package ecrypto_test

import (
	"reflect"
	"testing"

	"github.com/epes/ecrypto"
)

func Test_HexToBytes(t *testing.T) {
	tests := []struct {
		h    string
		want []byte
		err  bool
	}{
		{
			h:    "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			want: []byte{73, 39, 109, 32, 107, 105, 108, 108, 105, 110, 103, 32, 121, 111, 117, 114, 32, 98, 114, 97, 105, 110, 32, 108, 105, 107, 101, 32, 97, 32, 112, 111, 105, 115, 111, 110, 111, 117, 115, 32, 109, 117, 115, 104, 114, 111, 111, 109},
			err:  false,
		},
		{
			h:    "abc",
			want: nil,
			err:  true,
		},
	}

	for _, tc := range tests {
		got, err := ecrypto.HexToBytes(tc.h)

		if tc.err {
			if err == nil {
				t.Fatalf("should have gotten an error. in: %v, want: %v", tc.h, tc.want)
			}
		} else {
			if err != nil {
				t.Fatalf("encountered unexpected error: %s", err.Error())
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		}
	}
}

func Test_ByteToHex(t *testing.T) {
	tests := []struct {
		b byte
		h string
	}{
		{
			b: 0,
			h: "00",
		},
		{
			b: 255,
			h: "ff",
		},
	}

	for _, tc := range tests {
		got := ecrypto.ByteToHex(tc.b)

		if got != tc.h {
			t.Fatalf("expected: %v, got: %v", tc.h, got)
		}
	}
}
