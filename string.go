package ecrypto

func StringsToBytes(ss []string) [][]byte {
	b := make([][]byte, len(ss))

	for i, s := range ss {
		b[i] = []byte(s)
	}

	return b
}

func InferStringsToBytes(ss []string) [][]byte {
	b := make([][]byte, len(ss))

	for i, s := range ss {
		b[i] = InferEncoding(s)
	}

	return b
}

func HexStringsToBytes(ss []string) [][]byte {
	b := make([][]byte, len(ss))

	for i, s := range ss {
		b[i] = MustHexToBytes(s)
	}

	return b
}

func B64StringsToBytes(ss []string) [][]byte {
	b := make([][]byte, len(ss))

	for i, s := range ss {
		b[i] = MustB64ToBytes(s)
	}

	return b
}
