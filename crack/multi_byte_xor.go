package crack

import (
	"math"

	"github.com/epes/ecrypto"
	"github.com/epes/ecrypto/xor"
)

func MultiByteXOR(b []byte) (result []byte, key []byte, weight int) {
	return MultiByteXORKeyLengthRange(b, 1, 40)
}

func MultiByteXORKeyLengthRange(b []byte, min int, max int) (result []byte, key []byte, weight int) {
	keysizes := ecrypto.HammingNLowestBlockSizesInRange(b, 5, min, max)

	weight = math.MinInt32

	for _, ks := range keysizes {
		bbs := ecrypto.BytesBlockSplit(b, ks)
		tbs := ecrypto.Transpose(bbs)

		kkey := make([]byte, ks)
		var kweight int

		for bindex, block := range tbs {
			_, k, w := OneByteXOR(block)

			kkey[bindex] = k[0]
			kweight += w
		}

		if kweight > weight {
			weight = kweight
			key = kkey
		}
	}

	result = xor.DecryptECB(b, key)

	return
}
