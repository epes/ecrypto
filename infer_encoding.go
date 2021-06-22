package ecrypto

import "fmt"

func InferEncoding(s string) []byte {
	bits := true
	hex := true
	b64 := true

	b := []byte(s)

	if len(b)%2 != 0 {
		// default to ascii
		fmt.Println("inferred input to be ASCII")
		return []byte(s)
	}

	if len(b)%4 != 0 {
		b64 = false
	}

	if len(b)%8 != 0 {
		bits = false
	}

	for _, c := range b {
		if bits && !IsBit(c) {
			bits = false
		}

		if hex && !IsHex(c) {
			hex = false
		}

		if b64 && !IsB64(c) {
			b64 = false
		}

		if Trues(bits, hex, b64) == 0 {
			break
		}
	}

	if Trues(bits, hex, b64) == 0 {
		// default to ascii
		fmt.Println("inferred input to be ASCII")
		return []byte(s)
	}

	// if it could be one or more, choose the most restrictive
	if Trues(bits, hex, b64) > 0 {
		if bits {
			hex = false
			b64 = false
		} else if hex {
			b64 = false
		}
	}

	if bits {
		fmt.Println("inferred input to be bits")
		return MustBitsToBytes(s)
	}

	if hex {
		fmt.Println("inferred input to be hex")
		return MustHexToBytes(s)
	}

	if b64 {
		fmt.Println("inferred input to be b64")
		return MustB64ToBytes(s)
	}

	panic("it should have never come to this in infer")
}
