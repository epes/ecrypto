package tls

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/epes/ecrypto"

	"github.com/epes/emath"
)

var (
	TypeApplicationData = byte(23)

	Version12 = []byte{3, 3}
)

type TLS struct {
	t         byte
	version   []byte
	dataSize  int
	iv        []byte
	encrypted []byte
	decrypted []byte
	decryptor cipher.BlockMode
}

func Decode(payload []byte, key []byte) (*TLS, error) {
	if len(payload) < 5 {
		return nil, fmt.Errorf("payload smaller than 5")
	}

	dataSize := int(emath.BytesToUint16(payload[3:5]))

	if len(payload) != 5+dataSize {
		return nil, fmt.Errorf("data split across multiple payloads. payload size: %d, dataSize: %d, payload: %v", len(payload), 5+dataSize, payload)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	decoded := &TLS{
		t:        payload[0],
		version:  payload[1:3],
		dataSize: dataSize,
	}

	if dataSize >= 16 {
		decoded.iv = payload[5:21]
		decoded.decryptor = cipher.NewCBCDecrypter(block, decoded.iv)
	}

	decoded.encrypted = payload[21:]

	return decoded, nil
}

func (t *TLS) String() string {
	return fmt.Sprintf("Type: %s | Version: %s | Data Size: %s | IV: %s\nEncrypted: %s\nDecrypted Bytes: %s\nDecrypted: %s\n", t.Type(), t.Version(), t.DataSize(), t.IV(), t.Encrypted(), ecrypto.BytesToHex(t.Decrypted()), t.Decrypted())
}

func (t *TLS) Type() string {
	switch t.t {
	case TypeApplicationData:
		return "Application Data"
	default:
		return "unknown"
	}
}

func (t *TLS) Version() string {
	if bytes.Equal(t.version, Version12) {
		return "1.2"
	}

	return "unknown"
}

func (t *TLS) DataSize() string {
	return strconv.Itoa(t.dataSize)
}

func (t *TLS) IV() string {
	if len(t.iv) == 0 {
		return "empty"
	}

	return hex.EncodeToString(t.iv)
}

func (t *TLS) Encrypted() string {
	return hex.EncodeToString(t.encrypted)
}

func (t *TLS) Decrypted() []byte {
	if t.decrypted != nil {
		return t.decrypted
	}

	result := make([]byte, len(t.encrypted))

	t.decryptor.CryptBlocks(result, t.encrypted)
	t.decrypted = result

	return t.decrypted
}
