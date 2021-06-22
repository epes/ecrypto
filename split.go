package ecrypto

// BytesBlockSplit splits bytes into k-sized byte slices + leftover slice.
func BytesBlockSplit(b []byte, k int) [][]byte {
	var bl [][]byte

	var last []byte

	if len(b)%k != 0 {
		last = b[len(b)/k*k:]
	}

	for i := 0; i < len(b)/k*k; i += k {
		bl = append(bl, b[i:i+k])
	}

	if last != nil {
		bl = append(bl, last)
	}

	return bl
}

func Transpose(bb [][]byte) [][]byte {
	if len(bb) == 0 {
		return nil
	}

	l := len(bb[0])
	if l == 0 {
		return nil
	}

	t := make([][]byte, l)

	for _, block := range bb {
		if len(block) > l {
			// TODO: more validation for odd shape matrix
			// oddly shaped matrix
			return nil
		}
	}

	for i := 0; i < l; i++ {
		t[i] = make([]byte, 0, len(bb))
	}

	for _, block := range bb {
		for i, b := range block {
			t[i] = append(t[i], b)
		}
	}

	return t
}

// RepeatedBlocks takes a byte slice and a block size and returns
// the number of duplicate blocks within the slice.
// [A] [A] [A] = 2 repeated blocks (1 A is original)
// [A] [B] [A] [B] [A] = 3 repeated blocks (2 As and 1 B)
func RepeatedBlocks(b []byte, k int) int {
	cache := make(map[string]int)

	bbs := BytesBlockSplit(b, k)

	for _, bb := range bbs {
		cache[BytesToHex(bb)]++
	}

	var repeats int

	for _, r := range cache {
		repeats += r - 1
	}

	return repeats
}

func MostRepeatedBlocks(bb [][]byte, k int) (mostslice []byte, repeats int) {
	for _, byteslice := range bb {
		crepeats := RepeatedBlocks(byteslice, k)

		if crepeats > repeats {
			mostslice = byteslice
			repeats = crepeats
		}
	}

	return
}
