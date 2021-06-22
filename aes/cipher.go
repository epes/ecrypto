package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type c struct {
	key []byte

	// TODO: implement own version
	goaes cipher.Block
}

func NewCipher(key []byte) (cipher.Block, error) {
	size := len(key)

	if !(size == 16 || size == 24 || size == 32) {
		return nil, fmt.Errorf("AES block size not 16, 24, or 32. got: %d", size)
	}

	goaes, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create aes cipher: %w", err)
	}

	return &c{key, goaes}, nil
}

func (a *c) BlockSize() int {
	return a.goaes.BlockSize()
}

func (a *c) Encrypt(dst []byte, src []byte) {
	a.goaes.Encrypt(dst, src)
}

func (a *c) Decrypt(dst []byte, src []byte) {
	a.goaes.Decrypt(dst, src)
}
