package ecrypto_test

import (
	"reflect"
	"testing"

	"github.com/epes/ecrypto"
)

func Test_BitsToBytes(t *testing.T) {
	tests := []struct {
		i    string
		want []byte
		err  bool
	}{
		{
			i:    "0000111100001111",
			want: []byte{15, 15},
			err:  false,
		},
		{
			i:    "1111111111111111",
			want: []byte{255, 255},
			err:  false,
		},
		{
			i:    "1001",
			want: nil,
			err:  true,
		},
	}

	for _, tc := range tests {
		got, err := ecrypto.BitsToBytes(tc.i)

		if tc.err {
			if err == nil {
				t.Fatalf("should have gotten an error. in: %v, want: %v", tc.i, tc.want)
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

func Test_BitsToByte(t *testing.T) {
	tests := []struct {
		b    string
		want byte
		err  bool
	}{
		{
			b:    "00000000",
			want: 0,
			err:  false,
		},
		{
			b:    "00000001",
			want: 1,
			err:  false,
		},
		{
			b:    "11111111",
			want: 255,
			err:  false,
		},
		{
			b:    "10101010",
			want: 170,
			err:  false,
		},
		{
			b:    "00000002",
			want: 0,
			err:  true,
		},
		{
			b:    "000",
			want: 0,
			err:  true,
		},
	}

	for _, tc := range tests {
		got, err := ecrypto.BitOctetToByte([]byte((tc.b)))

		if tc.err {
			if err == nil {
				t.Fatalf("should have gotten an error. in: %v, want: %v", tc.b, tc.want)
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
