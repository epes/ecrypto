package crack

import (
	"math"

	"github.com/epes/ecrypto/lang"
	"github.com/epes/ecrypto/xor"
)

func OneByteXOR(b []byte) (result []byte, key []byte, weight int) {
	weight = math.MinInt32

	for i := byte(0); ; i++ {
		ckey := []byte{i}
		cresult := xor.EncryptECB(b, ckey)
		cweight := lang.EnglishWeight(cresult)

		if cweight > weight {
			result = cresult
			key = ckey
			weight = cweight
		}

		if i == 255 {
			break
		}
	}

	return
}

func OneOfOneByteXOR(bb [][]byte) (result []byte, key []byte, weight int) {
	weight = math.MinInt32

	for _, b := range bb {
		cresult, ckey, cweight := OneByteXOR(b)

		if cweight > weight {
			result = cresult
			key = ckey
			weight = cweight
		}
	}

	return
}
