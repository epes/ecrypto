package ecrypto_test

import (
	"reflect"
	"testing"

	"github.com/epes/ecrypto"
)

func Test_InferInput(t *testing.T) {
	tests := []struct {
		i    string
		want []byte
	}{
		{
			i:    "ff",
			want: []byte{255},
		},
		{
			i:    "fff",
			want: []byte{102, 102, 102},
		},
		{
			i:    "fffg",
			want: []byte{125, 247, 224},
		},
		{
			i:    "00001111",
			want: []byte{15},
		},
	}

	for _, tc := range tests {
		got := ecrypto.InferEncoding(tc.i)

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
