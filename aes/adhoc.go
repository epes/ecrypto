package aes

import "fmt"

func DecryptECB(src []byte, key []byte) ([]byte, error) {
	cipher, err := NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	dst := make([]byte, len(src))

	for i := 0; i < len(src); i += cipher.BlockSize() {
		cipher.Decrypt(dst[i:i+cipher.BlockSize()], src[i:i+cipher.BlockSize()])
	}

	return dst, nil
}

func MustDecryptECB(src []byte, key []byte) []byte {
	m, err := DecryptECB(src, key)
	if err != nil {
		panic(err)
	}

	return m
}

func EncryptECB(src []byte, key []byte) ([]byte, error) {
	cipher, err := NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	dst := make([]byte, len(src))

	for i := 0; i < len(src); i += cipher.BlockSize() {
		cipher.Encrypt(dst[i:i+cipher.BlockSize()], src[i:i+cipher.BlockSize()])
	}

	return dst, nil
}
