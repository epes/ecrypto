package set_1

import (
	"testing"

	"github.com/epes/ecrypto"
	"github.com/epes/ecrypto/xor"
)

func Test_FixedXOR(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{
			a:    "1c0111001f010100061a024b53535009181c",
			b:    "686974207468652062756c6c277320657965",
			want: "746865206b696420646f6e277420706c6179",
		},
	}

	for _, tc := range tests {
		got := ecrypto.BytesToHex(xor.EncryptECB(
			ecrypto.MustHexToBytes(tc.a),
			ecrypto.MustHexToBytes(tc.b),
		))

		if tc.want != got {
			t.Fatalf("expected: %s, got: %s", tc.want, got)
		}
	}
}
