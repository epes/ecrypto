package crack

import "math"

type cracker func([]byte) ([]byte, []byte, int)

// TODO(epes): need better heuristic to determine English-ness
// of message before this works accurately
var crackers = []cracker{
	OneByteXOR,
	MultiByteXOR,
}

func Crack(b []byte) []byte {
	var result []byte
	var weight = math.MinInt32

	for _, c := range crackers {
		cresult, _, cweight := c(b)

		if cweight > weight {
			result = cresult
			weight = cweight
		}
	}

	return result
}

func OneOfCrack(bb [][]byte) []byte {
	var result []byte
	var weight = math.MinInt32

	for _, b := range bb {
		for _, c := range crackers {
			cresult, _, cweight := c(b)

			if cweight > weight {
				result = cresult
				weight = cweight
			}
		}
	}

	return result
}
