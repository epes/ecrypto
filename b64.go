package ecrypto

import "fmt"

const (
	b64encode = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var (
	b64decode = map[byte]byte{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8,
		'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15, 'Q': 16, 'R': 17,
		'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23, 'Y': 24, 'Z': 25,
		'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31, 'g': 32, 'h': 33, 'i': 34,
		'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39, 'o': 40, 'p': 41, 'q': 42, 'r': 43,
		's': 44, 't': 45, 'u': 46, 'v': 47, 'w': 48, 'x': 49, 'y': 50, 'z': 51,
		'0': 52, '1': 53, '2': 54, '3': 55, '4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61,
		'+': 62, '/': 63, '=': 0,
	}
)

func B64ToBytes(b64 string) ([]byte, error) {
	b64b := []byte(b64)

	if len(b64b)%4 != 0 {
		return nil, fmt.Errorf("b64 input not divisible by 4: %s - %v", b64, b64b)
	}

	b := make([]byte, 0, len(b64b)/4*3)

	for i := 0; i < len(b64b); i += 4 {
		qb, err := B64QuadToTripleBytes([4]byte{b64b[i], b64b[i+1], b64b[i+2], b64b[i+3]})
		if err != nil {
			return nil, fmt.Errorf("failed to convert b64 to byte: %w", err)
		}

		b = append(b, qb[0])

		if b64b[i+2] != '=' {
			b = append(b, qb[1])
		}

		if b64b[i+3] != '=' {
			b = append(b, qb[2])
		}
	}

	return b, nil
}

func MustB64ToBytes(b64 string) []byte {
	b, err := B64ToBytes(b64)
	if err != nil {
		panic(err)
	}

	return b
}

func B64QuadToTripleBytes(b [4]byte) ([3]byte, error) {
	one, ok1 := b64decode[b[0]]
	two, ok2 := b64decode[b[1]]
	three, ok3 := b64decode[b[2]]
	four, ok4 := b64decode[b[3]]

	if !(ok1 && ok2 && ok3 && ok4) {
		return [3]byte{}, fmt.Errorf("b64 quad contains non-b64 characters: %v", b)
	}

	return [3]byte{
		(one << 2) + (two >> 4),
		(two << 4) + (three >> 2),
		(three << 6) + four,
	}, nil
}

func BytesToB64(b []byte) string {
	var eqs int

	if len(b)%3 == 1 {
		b = append(b, 0, 0)
		eqs = 2
	} else if len(b)%3 == 2 {
		b = append(b, 0)
		eqs = 1
	}

	s := make([]byte, 0, len(b)/3)

	for i := 0; i < len(b); i += 3 {
		s = append(s, ByteTripleToB64([3]byte{b[i], b[i+1], b[i+2]})...)
	}

	for i := 0; i < eqs; i++ {
		s[len(s)-1-i] = '='
	}

	return string(s)
}

func ByteTripleToB64(b [3]byte) []byte {
	one := b[0] >> 2
	two := (b[0] << 4 & 63) + (b[1] >> 4)
	three := (b[1] << 2 & 63) + (b[2] >> 6)
	four := b[2] & 63

	return []byte{
		b64encode[one],
		b64encode[two],
		b64encode[three],
		b64encode[four],
	}
}

func IsB64(b byte) bool {
	_, ok := b64decode[b]
	return ok
}
