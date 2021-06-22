package set_1

import (
	"testing"

	"github.com/epes/ecrypto"
	"github.com/epes/ecrypto/xor"
)

func Test_RepeatingKeyXOR(t *testing.T) {
	tests := []struct {
		msg  string
		key  string
		want string
	}{
		{
			msg:  "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal",
			key:  "ICE",
			want: "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f",
		},
	}

	for _, tc := range tests {
		got := ecrypto.BytesToHex(xor.EncryptECB([]byte(tc.msg), []byte(tc.key)))

		if tc.want != got {
			t.Fatalf("expected: %s, got: %s", tc.want, got)
		}
	}
}
