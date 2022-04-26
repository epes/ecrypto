package tls_test

import (
	"testing"

	"github.com/epes/ecrypto"
	"github.com/epes/ecrypto/tls"
)

func Test_DecodeTLS(t *testing.T) {
	tests := []struct {
		packet string
		key    string
	}{
		{
			packet: "17 03 03 00 30 61 62 63 64 65 66 67 68 69 6a 6b 6c 6d 6e 6f 70 97 83 48 8a f5 fa 20 bf 7a 2e f6 9d eb b5 34 db 9f b0 7a 8c 27 21 de e5 40 9f 77 af 0c 3d de 56",
			key:    "752a18e7a9fcb7cbcdd8f98dd8f769eb",
		},
		{
			packet: "17 03 03 00 30 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 6c 42 1c 71 c4 2b 18 3b fa 06 19 5d 13 3d 0a 09 d0 0f c7 cb 4e 0f 5d 1c da 59 d1 47 ec 79 0c 99",
			key:    "f656d037b173ef3e11169f27231a84b6",
		},
	}

	for _, tc := range tests {
		_, err := tls.Decode(ecrypto.InferEncoding(tc.packet), ecrypto.InferEncoding(tc.key))
		if err != nil {
			t.Error(err)
		}
	}
}
