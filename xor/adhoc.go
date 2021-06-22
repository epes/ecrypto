package xor

func EncryptECB(src []byte, key []byte) []byte {
	cipher := NewCipher(key)

	dst := make([]byte, len(src))

	short := len(src) % cipher.BlockSize()
	fblock := len(src) / cipher.BlockSize() * cipher.BlockSize()

	for i := 0; i < fblock; i += cipher.BlockSize() {
		cipher.Encrypt(dst[i:i+cipher.BlockSize()], src[i:i+cipher.BlockSize()])
	}

	if short > 0 {
		cipher.Encrypt(dst[fblock:fblock+short], src[fblock:fblock+short])
	}

	return dst
}

func DecryptECB(src []byte, key []byte) []byte {
	return EncryptECB(src, key)
}
