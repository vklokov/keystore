package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
)

const (
	SECRET_LENGTH       = 32
	CYPHER_BLOCK_LENGTH = 16
)

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(str string) []byte {
	data, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		panic(err)
	}

	return data
}

func EncryptRecord(content, secret string) (string, error) {
	bytes := []byte(os.Getenv("CYPHER_BYTES")[0:CYPHER_BLOCK_LENGTH])
	chunk := secret[0:SECRET_LENGTH]
	block, err := aes.NewCipher([]byte(chunk))

	if err != nil {
		return "", err
	}

	raw := []byte(content)
	cf := cipher.NewCFBEncrypter(block, bytes)
	cfText := make([]byte, len(raw))
	cf.XORKeyStream(cfText, raw)

	return Base64Encode(cfText), nil
}

func DecryptRecord(content, secret string) (string, error) {
	bytes := []byte(os.Getenv("CYPHER_BYTES")[0:CYPHER_BLOCK_LENGTH])
	chunk := secret[0:SECRET_LENGTH]
	block, err := aes.NewCipher([]byte(chunk))

	if err != nil {
		return "", err
	}

	cf := cipher.NewCFBDecrypter(block, bytes)
	cfText := Base64Decode(content)
	raw := make([]byte, len(cfText))
	cf.XORKeyStream(raw, cfText)

	return string(raw), nil
}
