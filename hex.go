package ecrypto

import "fmt"

const (
	b16encode = "0123456789abcdef"
)

var (
	b16decode = map[byte]byte{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9,
		'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14, 'f': 15,
		'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15,
	}
)

func HexToBytes(h string) ([]byte, error) {
	hb := []byte(h)

	if len(hb)%2 != 0 {
		return nil, fmt.Errorf("hex input not divisible by 2: %s - %v", h, hb)
	}

	b := make([]byte, 0, len(hb)/2)

	for i := 0; i < len(hb); i += 2 {
		tb, err := HexTupleToByte([2]byte{hb[i], hb[i+1]})
		if err != nil {
			return nil, fmt.Errorf("failed to convert hex to byte: %w", err)
		}

		b = append(b, tb)
	}

	return b, nil
}

func MustHexToBytes(h string) []byte {
	b, err := HexToBytes(h)
	if err != nil {
		panic(err)
	}

	return b
}

func HexTupleToByte(h [2]byte) (byte, error) {
	h0, ok := b16decode[h[0]]
	if !ok {
		return 0, fmt.Errorf("hex tuple contains non-hex characters: %v - %v", h[0], h)
	}

	h1, ok := b16decode[h[1]]
	if !ok {
		return 0, fmt.Errorf("hex tuple contains non-hex characters: %v - %v", h[1], h)
	}

	return (h0 << 4) + h1, nil
}

func BytesToHex(b []byte) string {
	h := make([]byte, 0, len(b)*2)

	for _, c := range b {
		h = append(h, b16encode[c>>4])
		h = append(h, b16encode[c&15])
	}

	return string(h)
}

func ByteToHex(b byte) string {
	l := b >> 4
	r := b & 15

	return string([]byte{b16encode[l], b16encode[r]})
}

func IsHex(b byte) bool {
	_, ok := b16decode[b]
	return ok
}
