package xor

import "crypto/cipher"

type c struct {
	key []byte
}

func NewCipher(key []byte) cipher.Block {
	return &c{key: key}
}

func (x *c) BlockSize() int {
	return len(x.key)
}

func (x *c) Encrypt(dst []byte, src []byte) {
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] ^ x.key[i%x.BlockSize()]
	}
}

func (x *c) Decrypt(dst []byte, src []byte) {
	x.Encrypt(dst, src)
}
