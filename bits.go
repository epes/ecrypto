package ecrypto

import "fmt"

func BitsToBytes(s string) ([]byte, error) {
	sb := []byte(s)

	if len(sb)%8 != 0 {
		return nil, fmt.Errorf("bits input not divisible by 8: %s - %v", s, sb)
	}

	b := make([]byte, 0, len(sb)/8)

	for i := 0; i < len(sb); i += 8 {
		ob, err := BitOctetToByte(sb[i : i+8])
		if err != nil {
			return nil, fmt.Errorf("failed to convert bits to byte: %w", err)
		}

		b = append(b, ob)
	}

	return b, nil
}

func MustBitsToBytes(s string) []byte {
	b, err := BitsToBytes(s)
	if err != nil {
		panic(err)
	}

	return b
}

func BitOctetToByte(s []byte) (byte, error) {
	if len(s) != 8 {
		return 0, fmt.Errorf("bits length is not 8: %v", s)
	}

	var b byte

	for i := 0; i < 8; i++ {
		if IsBit(s[i]) {
			b <<= 1
			b += s[i] - 48
		} else {
			return 0, fmt.Errorf("bits contain non-binary characters: %v - %s", s[i], s)
		}
	}

	return b, nil
}

func IsBit(b byte) bool {
	return b == '0' || b == '1'
}

func CountTrueBits(b byte) int {
	var h byte

	for ; b > 0; b >>= 1 {
		h += b & 1
	}

	return int(h)
}
