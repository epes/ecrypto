package lang

var (
	enMap = map[byte]int{
		'0': 1, '1': 1, '2': 1, '3': 1, '4': 1, '5': 1, '6': 1, '7': 1, '8': 1, '9': 1,
		'a': 40, 'A': 40, 'b': 8, 'B': 8, 'c': 15, 'C': 15, 'd': 22, 'D': 22, 'e': 60, 'E': 60,
		'f': 13, 'F': 13, 'g': 8, 'G': 8, 'h': 32, 'H': 32, 'i': 40, 'I': 40, 'j': 2, 'J': 2,
		'k': 4, 'K': 4, 'l': 20, 'L': 20, 'm': 15, 'M': 15, 'n': 40, 'N': 40, 'o': 40, 'O': 40,
		'p': 8, 'P': 8, 'q': 2, 'Q': 2, 'r': 31, 'R': 31, 's': 40, 'S': 40, 't': 45, 'T': 45,
		'u': 17, 'U': 17, 'v': 6, 'V': 6, 'w': 10, 'W': 10, 'x': 2, 'X': 2, 'y': 10, 'Y': 10,
		'z': 1, 'Z': 1, '.': 1, '!': 1, '?': 1, 32: 10,
	}
)

// TODO: need better heuristic for english words
func EnglishWeight(m []byte) int {
	var weight int

	for _, b := range m {
		if w, ok := enMap[b]; ok {
			weight += w
		}
	}

	return weight
}
